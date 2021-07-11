// Copyright (c) Inlets Author(s) 2021. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package main

import (
	"fmt"
	"os"

	"github.com/inlets/inlets/cmd"
	"github.com/sirupsen/logrus"
)

// These values will be injected into these variables at the build time.
var (
	Version   string
	GitCommit string
)

func main() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006/01/02 15:04:05"
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	if err := cmd.Execute(Version, GitCommit); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
