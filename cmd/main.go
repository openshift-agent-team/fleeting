package main

import "github.com/openshift-agent-team/fleeting/pkg/isosource"

func main() {
	_, err := isosource.EnsureIso()
	if err != nil {
		panic(err)
	}
}
