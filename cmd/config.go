package cmd

import (
	"fmt"
	"os"

	"github.com/YoshinoriN/cahsper-cli/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(configCommand)
	configCommand.AddCommand(showConfigCommmand)
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
		fmt.Println("CAHSPER_PASSWORD:", os.Getenv("CAHSPER_PASSWORD"))
		fmt.Println("CAHSPER_USER_POOL_ID:", os.Getenv("CAHSPER_USER_POOL_ID"))
		fmt.Println("CAHSPER_AWS_COGNITO_APP_CLIENT_ID:", os.Getenv("CAHSPER_AWS_COGNITO_APP_CLIENT_ID"))
		fmt.Println("CAHSPER_SERVER_URL:", os.Getenv("CAHSPER_SERVER_URL"))
		fmt.Println("CAHSPER_ID_TOKEN:", os.Getenv("CAHSPER_ID_TOKEN"))
		fmt.Println("CAHSPER_ACCESS_TOKEN:", os.Getenv("CAHSPER_ACCESS_TOKEN"))
	},
}
