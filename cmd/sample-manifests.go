package cmd

import (
	"fmt"
	"os"

	"github.com/openshift-agent-team/fleeting/pkg/manifests"
	"github.com/spf13/cobra"
)

var manifestPath string
var sampleManifests = &cobra.Command{
	Use:   "sample-manifests",
	Short: "Generates example manifests for a cluster deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		argPath, err := cmd.Flags().GetString("manifest-dir")
		if err != nil {
			return nil
		} else {
			if argPath != "" {
				dirInfo, err := os.Stat(argPath)
				if os.IsNotExist(err) {
					fmt.Fprintf(os.Stderr, "The path to put the manifests '%s' does not exist\n", argPath)
					os.Exit(1)
				} else if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if !dirInfo.IsDir() {
					fmt.Fprintf(os.Stderr, "%s is not a directory\n", argPath)
				}
			}
		}
		err = manifests.GenerateManifests(argPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failure encountered in manifest generation. Error: %s", err)
			os.Exit(1)
		}
		return err
	},
}

func init() {
	sampleManifests.Flags().StringVarP(&manifestPath, "manifest-dir", "o", "", "Path to put the generated manifests. If empty, prints to stdout")
	rootCmd.AddCommand(sampleManifests)
}
