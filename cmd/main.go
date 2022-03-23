package main

import (
	"flag"
	"os"

	"github.com/openshift-agent-team/fleeting/pkg/imagebuilder"
	"github.com/openshift-agent-team/fleeting/pkg/isosource"
)

func main() {
	nodeZeroIP := flag.String("node-zero-ip", "", "IP of the node to run OpenShift Assisted Installation Service on. (Required)")
	apiVip := flag.String("apivip", "", "API Virtual IP. (Required)")
	flag.Parse()

	if *nodeZeroIP == "" || *apiVip == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	baseImage, err := isosource.EnsureIso()
	if err != nil {
		panic(err)
	}

	err = imagebuilder.BuildImage(baseImage, *nodeZeroIP, *apiVip)
	if err != nil {
		panic(err)
	}
}
