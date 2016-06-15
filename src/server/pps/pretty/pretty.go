package pretty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"text/template"

	"github.com/Jeffail/gabs"
	"github.com/fatih/color"
	ppsclient "github.com/pachyderm/pachyderm/src/client/pps"
)

func PrintJobHeader(w io.Writer) {
	fmt.Fprint(w, "ID\tOUTPUT\tSTATE\t\n")
}

func PrintJobInfo(w io.Writer, jobInfo *ppsclient.JobInfo) {
	fmt.Fprintf(w, "%s\t", jobInfo.Job.ID)
	if jobInfo.OutputCommit != nil {
		fmt.Fprintf(w, "%s/%s\t", jobInfo.OutputCommit.Repo.Name, jobInfo.OutputCommit.ID)
	} else {
		fmt.Fprintf(w, "-\t")
	}
	fmt.Fprintf(w, "%s\t\n", jobState(jobInfo))
}

func PrintPipelineHeader(w io.Writer) {
	fmt.Fprint(w, "NAME\tINPUT\tOUTPUT\tSTATE\t\n")
}

func PrintPipelineInfo(w io.Writer, pipelineInfo *ppsclient.PipelineInfo) {
	fmt.Fprintf(w, "%s\t", pipelineInfo.Pipeline.Name)
	if len(pipelineInfo.Inputs) == 0 {
		fmt.Fprintf(w, "\t")
	} else {
		for i, input := range pipelineInfo.Inputs {
			fmt.Fprintf(w, "%s", input.Repo.Name)
			if i == len(pipelineInfo.Inputs)-1 {
				fmt.Fprintf(w, "\t")
			} else {
				fmt.Fprintf(w, ", ")
			}
		}
	}
	if pipelineInfo.OutputRepo != nil {
		fmt.Fprintf(w, "%s\t", pipelineInfo.OutputRepo.Name)
	} else {
		fmt.Fprintf(w, "\t")
	}
	fmt.Fprintf(w, "%s\t\n", pipelineState(pipelineInfo))
}

func PrintPipelineInputHeader(w io.Writer) {
	fmt.Fprint(w, "NAME\tPARTITION\tINCREMENTAL\t\n")
}

func PrintPipelineInput(w io.Writer, pipelineInput *ppsclient.PipelineInput) {
	fmt.Fprintf(w, "%s\t", pipelineInput.Repo.Name)
	fmt.Fprintf(w, "%s\t", pipelineInput.Method.Partition)
	fmt.Fprintf(w, "%t\t\n", pipelineInput.Method.Incremental)
}

func pipelineInputs(pipelineInfo *ppsclient.PipelineInfo) string {
	var buffer bytes.Buffer
	writer := tabwriter.NewWriter(&buffer, 20, 1, 3, ' ', 0)
	PrintPipelineInputHeader(writer)
	for _, input := range pipelineInfo.Inputs {
		PrintPipelineInput(writer, input)
	}
	// can't error because buffer can't error on Write
	writer.Flush()
	return buffer.String()
}

func PrintDetailedJobInfo(jobInfo *ppsclient.JobInfo) {
	bytes, err := json.Marshal(jobInfo)
	if err != nil {
		fmt.Println(err.Error())
	}

	obj, err := gabs.ParseJSON(bytes)
	if err != nil {
		fmt.Println(err.Error())
	}

	// state is an integer; we want to print a string
	_, err = obj.Set(ppsclient.JobState_name[int32(jobInfo.State)], "state")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(obj.StringIndent("", "    "))
}

var funcMap template.FuncMap = template.FuncMap{
	"pipelineState":  pipelineState,
	"pipelineInputs": pipelineInputs,
}

func PrintDetailedPipelineInfo(pipelineInfo *ppsclient.PipelineInfo) {
	template, err := template.New("PipelineInfo").Funcs(funcMap).Parse(
		`Name: {{.Pipeline.Name}}
Created: {{.CreatedAt}}
State: {{pipelineState .}}
Parallelism: {{.Parallelism}}
Inputs:
{{pipelineInputs .}}
Recent Error: {{.RecentError}}
`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = template.Execute(os.Stdout, pipelineInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func jobState(jobInfo *ppsclient.JobInfo) string {
	switch jobInfo.State {
	case ppsclient.JobState_JOB_PULLING:
		return color.New(color.FgYellow).SprintFunc()("pulling")
	case ppsclient.JobState_JOB_RUNNING:
		return color.New(color.FgYellow).SprintFunc()("running")
	case ppsclient.JobState_JOB_FAILURE:
		return color.New(color.FgRed).SprintFunc()("failure")
	case ppsclient.JobState_JOB_SUCCESS:
		return color.New(color.FgGreen).SprintFunc()("success")
	}
	return "-"
}

func pipelineState(pipelineInfo *ppsclient.PipelineInfo) string {
	switch pipelineInfo.State {
	case ppsclient.PipelineState_PIPELINE_STARTING:
		return color.New(color.FgYellow).SprintFunc()("starting")
	case ppsclient.PipelineState_PIPELINE_RUNNING:
		return color.New(color.FgGreen).SprintFunc()("running")
	case ppsclient.PipelineState_PIPELINE_RESTARTING:
		return color.New(color.FgYellow).SprintFunc()("restarting")
	case ppsclient.PipelineState_PIPELINE_FAILED:
		return color.New(color.FgRed).SprintFunc()("failed")
	}
	return "-"
}
