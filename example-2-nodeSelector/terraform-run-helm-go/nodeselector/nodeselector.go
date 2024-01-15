// Utility script to add nodeSelector values (provided as a JSON map via the first argument) to deployment YAML passed in via stdin, emitting the changed YAML to stderr.
// Intended for use as a helm post render step.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func marshal(input map[string]interface{}) {
	data, err := yaml.Marshal(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(string(data))
	fmt.Println("---")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a single JSON argument for nodeSelector values. If you provided one, maybe shell quoting issues?")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	decoder := yaml.NewDecoder(reader)

	for {
		var input map[string]interface{}
		err := decoder.Decode(&input)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		deploymentSpec, ok := input["spec"].(map[interface{}]interface{})
		if !ok {
			marshal(input)
			continue
			//return
		}

		template, ok := deploymentSpec["template"].(map[interface{}]interface{})
		if !ok {
			marshal(input)
			continue
		}
		podSpec, ok := template["spec"].(map[interface{}]interface{})
		if !ok {
			marshal(input)
			continue
		}

		var selectorsInput = os.Args[1]
		var selectors map[string]string
		yaml.Unmarshal([]byte(selectorsInput), &selectors)
		podSpec["nodeSelector"] = selectors
		marshal(input)
	}
}
