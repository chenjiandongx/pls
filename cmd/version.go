package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "pls version: [v0.1.4]"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the version of pls",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
	return cmd
}
