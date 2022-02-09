package main

import (
	"github.com/openshift-agent-team/fleeting/pkg/imagebuilder"
	"github.com/openshift-agent-team/fleeting/pkg/isosource"
)

func main() {
	baseImage, err := isosource.EnsureIso()
	if err != nil {
		panic(err)
	}

	err = imagebuilder.BuildImage(baseImage)
	if err != nil {
		panic(err)
	}
}
