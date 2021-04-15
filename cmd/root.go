package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "cahsper",
	Short: "CLI tool for Cahsper",
	Long: `Cahsper is an alternative of twitter for a solitary person.
                CLI: https://github.com/yoshinorin/cahsper-cli
                Server-side: https://github.com/yoshinorin/cahsper`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
