package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pachyderm/pachyderm/v2/src/internal/cmdutil"
	"github.com/pachyderm/pachyderm/v2/src/internal/errors"
	"github.com/pachyderm/pachyderm/v2/src/internal/log"
	"github.com/pachyderm/pachyderm/v2/src/server/cmd/pachctl/cmd"
	"github.com/pkg/errors"

	"github.com/spf13/cobra/doc"
)

type appEnv struct{}

func main() {
	log.InitPachctlLogger()
	cmdutil.Main(context.Background(), do, &appEnv{})
}

func do(ctx context.Context, appEnvObj interface{}) error {
	path := "./docs/"


	if err := os.MkdirAll(path, os.ModePerm); !os.IsExist(err) && err != nil {
		return errors.WithStack(fmt.Errorf("error creating directory: %w", err))
	}

	rootCmd, err := cmd.PachctlCmd()
	if err != nil {
		return errors.WithStack(fmt.Errorf("could not generate pachctl command: %w", err))

	}
	rootCmd.DisableAutoGenTag = true

	const fmTemplate = `---
date: %s
title: "%s"
slug: "Learn about the %s command"
---

`

	filePrepender := func(filename string) string {
		now := time.Now().Format(time.RFC3339)
		name := filepath.Base(filename)
		base := strings.TrimSuffix(name, filepath.Ext(name))
		return fmt.Sprintf(fmTemplate, now, strings.Replace(base, "_", " ", -1), base)
	}

	linkHandler := func(name string) string {
		base := strings.TrimSuffix(name, filepath.Ext(name))
		return "/commands/" + strings.ToLower(base) + "/"
	}

	err = doc.GenMarkdownTreeCustom(rootCmd, path, filePrepender, linkHandler)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to generate Markdown documentation: %w", err))

	}

	return nil
}
