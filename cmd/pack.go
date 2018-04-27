package cmd

import (
	"fmt"
	"os"

	"github.com/docker/lunchbox/internal"
	"github.com/docker/lunchbox/packager"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack <app-name> [-o output_file]",
	Short: "Pack this app as a single file",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := ""
		if len(args) > 0 {
			app = args[0]
		}
		err := packager.Pack(app, packOutputFile)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	},
}

var packOutputFile string

func init() {
	if internal.Experimental == "on" {
		rootCmd.AddCommand(packCmd)
		packCmd.Flags().StringVarP(&packOutputFile, "output", "o", "-", "Output file (- for stdout)")
	}
}
