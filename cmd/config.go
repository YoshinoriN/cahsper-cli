package cmd

import (
	"bufio"
	"fmt"
	"os"

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
		scanner := bufio.NewScanner(os.Stdin)
		config := utils.Config{}

		fmt.Print("Cognito AppClientId: ")
		scanner.Scan()
		config.Settings.Cognito.AppClientID = scanner.Text()

		fmt.Print("Cognito UserPoolID: ")
		scanner.Scan()
		config.Settings.Cognito.UserPoolID = scanner.Text()

		fmt.Print("ServerURL: ")
		scanner.Scan()
		config.Settings.ServerURL = scanner.Text()

		utils.Write(utils.GetConfigFilePath(), config)
	},
}
