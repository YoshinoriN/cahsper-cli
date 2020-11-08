package cmd

import (
	"github.com/YoshinoriN/cahsper-cli/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(configCommand)
	configCommand.AddCommand(showConfigCommmand)
	configCommand.AddCommand(setConfigCommand)
}

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "get/set your config for cahsper",
	Long:  "get/set your config for cahsper",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var showConfigCommmand = &cobra.Command{
	Use:   "list",
	Short: "Show all of cahsper configure variables.",
	Long:  `Show cahsper configure variables from config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: check file is already exists or not
		configFilePath := utils.GetConfigFilePath()
		config := utils.Read(configFilePath)
		utils.Print(config)
	},
}

// TODO: Refactor
var setConfigCommand = &cobra.Command{
	Use:   "set",
	Short: "Set cahsper configure variables.",
	Long:  `Set configure variables.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}
