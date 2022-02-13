package cmd

import (
	"fmt"

	"github.com/openshift-agent-team/fleeting/pkg/imagebuilder"
	"github.com/openshift-agent-team/fleeting/pkg/isosource"
	"github.com/spf13/cobra"
)

var buildIsoCmd = &cobra.Command{
	Use:   "build-iso",
	Short: "Builds an iso from the provided embeddable data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building rhcos image...")

		baseImage, err := isosource.EnsureIso()
		if err != nil {
			panic(err)
		}
		err = imagebuilder.BuildImage(baseImage)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildIsoCmd)
}
