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
package helm

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/ghodss/yaml"
)

type release struct {
	Name       string
	Revision   int32
	Updated    string
	Status     string
	Chart      string
	AppVersion string
	Namespace  string
}

type list struct {
	Next     string
	Releases []release
}

type helmStatus struct {
	status string
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
		os.Exit(-1)
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func parseOutput(outs []byte, helmstatus *helmStatus) {
	var status = regexp.MustCompile(`STATUS: (.*)`)
	result := status.FindStringSubmatch(string(outs))
	if len(result) > 0 {
		helmstatus.status = string(result[1])
	}
}

// Version ...
func Version() {
	fmt.Print("helm version: ")
	cmd := exec.Command("helm", "version", "--client", "--short")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
	output := cmdOutput.Bytes()
	printOutput(output)
}

// List ...
func List(namespace string) {
	cmd := exec.Command("helm", "list", "--namespace", namespace, "-c")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}

// ListAll ...
func ListAll() {
	cmd := exec.Command("helm", "list")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
	output := cmdOutput.Bytes()
	printOutput(output)
}

// ListSubCharts ...
func ListSubCharts(subShart string) {
	cmd := exec.Command("helm", "list", "--output", "yaml")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
	helmList := cmdOutput.Bytes()
	var p2 list
	err := yaml.Unmarshal(helmList, &p2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	for _, r := range p2.Releases {
		if r.Chart == subShart {
			Delete(r.Name, false)
		}
	}
}

// Delete chart
func Delete(chart string, dryRun bool) {
	var myargs []string
	if dryRun {
		myargs = []string{"helm", "delete", "--purge", chart, "--dry-run"}
	} else {
		myargs = []string{"delete", "--purge", chart}
	}
	cmd := exec.Command("helm", myargs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}

// GetHelmStatus ...
func GetHelmStatus(chart string) string {
	cmd := exec.Command("helm", "status", chart)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
	output := cmdOutput.Bytes()
	helmstatus := helmStatus{}
	parseOutput(output, &helmstatus)
	return helmstatus.status
}
