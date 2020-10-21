package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"-v", "--version"}, // Why does not work???
	Short:   "Show version number of cahsper-cli",
	Long:    "Show version number of cahsper-cli",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cahsper-cli v0.0.1")
	},
}
