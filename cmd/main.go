package main

import (
	"flag"
	"os"

	"github.com/openshift-agent-team/fleeting/pkg/imagebuilder"
	"github.com/openshift-agent-team/fleeting/pkg/isosource"
)

func main() {
	nodeZeroIP := flag.String("node-zero-ip", "", "IP of the node to run OpenShift Assisted Installation Service on. (Required)")
	flag.Parse()

	if *nodeZeroIP == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	baseImage, err := isosource.EnsureIso()
	if err != nil {
		panic(err)
	}

	err = imagebuilder.BuildImage(baseImage, *nodeZeroIP)
	if err != nil {
		panic(err)
	}
}
