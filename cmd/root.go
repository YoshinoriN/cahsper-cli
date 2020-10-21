package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cahsper",
	Short: "CLI tool for Cahsper",
	Long: `Cahsper is an alternative of twitter for a solitary person.
                CLI: https://github.com/YoshinoriN/cahsper-cli
                Server-side: https://github.com/YoshinoriN/cahsper`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	// TODO
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
