package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd *cobra.Command = &cobra.Command{
	Use:   "local-cli",
	Short: "testing local cli without remote container",
	Long:  "no specific use case but to experiment with cobra cli",
}

func init() {
	RootCmd.AddCommand(PrintNameCmd)
}
