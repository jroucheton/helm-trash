/*
(c) Copyright 2018, Gemalto. All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jroucheton/helm-trash/pkg/helm"

	"github.com/spf13/cobra"
)

type trashCmd struct {
	chartName    string
	chartVersion string
	namespace    string
	dryRun       bool
	debug        bool
}

var (
	globalUsage = `
This command delete sub charts which have been deployed through helm spray.

Arguments shall be the chart reference.

 $ helm trash umbrella-chart-1.0.0-rc.1+build.32

To check the generated manifests of a release without installing the chart,
the '--debug' and '--dry-run' flags can be combined. This will still require a
round-trip to the Tiller server.

To see the list of chart repositories, use 'helm repo list'. To search for
charts in a repository, use 'helm search'.
`
)

func newTrashCmd(args []string) *cobra.Command {

	p := &trashCmd{}

	cmd := &cobra.Command{
		Use:          "helm trash [CHART]",
		Short:        `Helm plugin to deletet releases previously installed through helm spray`,
		Long:         globalUsage,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) != 1 {
				return errors.New("This command needs at least 1 argument: chart name")
			}

			// TODO: check format for chart name (directory, url, tgz...)
			p.chartName = args[0]

			return p.trash()
		},
	}

	f := cmd.Flags()
	f.StringVarP(&p.namespace, "namespace", "n", "default", "namespace to target.")
	f.BoolVar(&p.dryRun, "dry-run", false, "simulate a trash")
	f.BoolVar(&p.debug, "debug", false, "enable verbose output")
	f.Parse(args)

	// When called through helm, debug mode is transmitted through the HELM_DEBUG envvar
	if !p.debug {
		if "1" == os.Getenv("HELM_DEBUG") {
			p.debug = true
		}
	}

	return cmd

}

func (p *trashCmd) trash() error {

	// For debug...
	if p.debug {
	}

	// List helm release
	helm.ListSubCharts(p.chartName)

	fmt.Println("[trash] trash completed.")

	return nil
}

func main() {
	cmd := newTrashCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
