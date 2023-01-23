package logging

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/fatih/camelcase"
	"github.com/pachyderm/pachyderm/v2/src/internal/errors"
	"github.com/pachyderm/pachyderm/v2/src/internal/log"
	"github.com/pachyderm/pachyderm/v2/src/internal/pctx"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

// This needs to be a global var, not a field on the logger, because multiple servers
// create new loggers, and the prometheus registration uses a global namespace
var reportDurationGauge prometheus.Gauge
var reportDurationsOnce sync.Once

type LoggingInterceptor struct {
	mutex     sync.Mutex // synchronizes access to both histogram and counter maps
	histogram map[string]*prometheus.HistogramVec
	counter   map[string]prometheus.Counter
	Level     log.Level
}

// NewLoggingInterceptor creates a new interceptor that logs method start and end.  Note that the
// provided context is only for warnings generated by this function.  It is not the root logger;
// that is injected by the BaseContextInterceptor.
func NewLoggingInterceptor(ctx context.Context) *LoggingInterceptor {
	interceptor := &LoggingInterceptor{
		histogram: make(map[string]*prometheus.HistogramVec),
		counter:   make(map[string]prometheus.Counter),
		Level:     log.InfoLevel,
	}

	reportDurationsOnce.Do(func() {
		newReportMetricGauge := prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "pachyderm",
				Subsystem: "pachd",
				Name:      "report_metric",
				Help:      "gauge of number of calls to reportDuration()",
			},
		)
		if err := prometheus.Register(newReportMetricGauge); err != nil {
			// metrics may be redundantly registered; ignore these errors
			if !errors.As(err, &prometheus.AlreadyRegisteredError{}) {
				log.Error(ctx, "error registering prometheus metric", zap.Error(err))
			}
		} else {
			reportDurationGauge = newReportMetricGauge
		}
	})
	return interceptor
}

func parseMethod(fullMethod string) (string, string) {
	fullMethod = strings.Trim(fullMethod, "/")
	parts := strings.SplitN(fullMethod, "/", 2)
	switch len(parts) {
	case 0:
		return "", ""
	case 1:
		return parts[0], ""
	default:
		return parts[0], parts[1]
	}
}

func (li *LoggingInterceptor) logUnaryAfter(ctx context.Context, lvl log.Level, service, method string, start time.Time, err error) {
	duration := time.Since(start)
	go li.reportDuration(service, method, duration, err)
	dolog(ctx, lvl, "response for "+service+"/"+method, zap.Duration("duration", duration))
}

func topLevelService(fullyQualifiedService string) string {
	tokens := strings.Split(fullyQualifiedService, ".")
	return tokens[0]
}

func (li *LoggingInterceptor) getHistogram(service string, method string) *prometheus.HistogramVec {
	fullyQualifiedName := fmt.Sprintf("%v/%v", service, method)
	histVec, ok := li.histogram[fullyQualifiedName]
	if !ok {
		var tokens []string
		for _, token := range camelcase.Split(method) {
			tokens = append(tokens, strings.ToLower(token))
		}
		rootStatName := strings.Join(tokens, "_")

		histogramName := fmt.Sprintf("%v_seconds", rootStatName)
		histVec = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "pachyderm",
				Subsystem: fmt.Sprintf("pachd_%v", topLevelService(service)),
				Name:      histogramName,
				Help:      fmt.Sprintf("Run time of %v", method),
				Buckets:   []float64{0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2, 5, 10, 30, 60, 120, 300, 600, 1800, 3600, 86400},
			},
			[]string{
				"state", // Since both finished and errored API calls can have run times
			},
		)
		if err := prometheus.Register(histVec); err != nil {
			// metrics may be redundantly registered; ignore these errors
			if !errors.As(err, &prometheus.AlreadyRegisteredError{}) {
				log.Info(pctx.TODO(), "error registering prometheus metric", zap.String("histogramName", histogramName), zap.Error(err))
			}
		} else {
			li.histogram[fullyQualifiedName] = histVec
		}
	}
	return histVec
}

func (li *LoggingInterceptor) reportDuration(service string, method string, duration time.Duration, err error) {
	// Count the number of reportDuration() goros in case we start to leak them
	if reportDurationGauge != nil {
		reportDurationGauge.Inc()
	}
	defer func() {
		if reportDurationGauge != nil {
			reportDurationGauge.Dec()
		}
	}()
	li.mutex.Lock() // for concurrent map access (histogram,counter)
	defer li.mutex.Unlock()

	state := "finished"
	if err != nil {
		state = "errored"
	}
	if hist, err := li.getHistogram(service, method).GetMetricWithLabelValues(state); err != nil {
		log.Info(pctx.TODO(), "failed to get histogram for state", zap.String("state", state), zap.Error(err))
	} else {
		hist.Observe(duration.Seconds())
	}
}
