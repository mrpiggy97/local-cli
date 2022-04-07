package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name *string = new(string)

func printName(cmd *cobra.Command, args []string) {
	fmt.Println(*name)
}

var PrintNameCmd *cobra.Command = &cobra.Command{
	Use:   "print-name",
	Short: "get name and print it",
	Run:   printName,
}

func init() {
	PrintNameCmd.Flags().StringVar(name, "name", "", "sets name to print")
	PrintNameCmd.MarkFlagRequired("name")
}
