package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(configCommand)
	configCommand.AddCommand(getConfigListCommmand)
}

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "get/set your config for cahsper",
	Long:  "get/set your config for cahsper",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var getConfigListCommmand = &cobra.Command{
	Use:   "list",
	Short: "Show all of cahsper configure variables.",
	Long: `Show cahsper configure variables.
	CAHSPER_USER_NAME: <your_aws_cognito_username>
	CAHSPER_PASSWORD: <your_aws_cognito_password>
	CAHSPER_USER_POOL_ID: <your_aws_cognito_userpool_id>
	CAHSPER_AWS_COGNITO_APP_CLIENT_ID: <your_aws_cognito_app_client_id>
	CAHSPER_SERVER_URL: <cahsper_server_side_url>
	CAHSPER_ID_TOKEN: <your_aws_cognito_idtoken_for_cahsper>
	CAHSPER_ACCESS_TOKEN: <your_aws_cognito_accesstoken_for_cahsper>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CAHSPER_USER_NAME:", os.Getenv("CAHSPER_USER_NAME"))
		fmt.Println("CAHSPER_PASSWORD:", os.Getenv("CAHSPER_PASSWORD"))
		fmt.Println("CAHSPER_USER_POOL_ID:", os.Getenv("CAHSPER_USER_POOL_ID"))
		fmt.Println("CAHSPER_AWS_COGNITO_APP_CLIENT_ID:", os.Getenv("CAHSPER_AWS_COGNITO_APP_CLIENT_ID"))
		fmt.Println("CAHSPER_SERVER_URL:", os.Getenv("CAHSPER_SERVER_URL"))
		fmt.Println("CAHSPER_ID_TOKEN:", os.Getenv("CAHSPER_ID_TOKEN"))
		fmt.Println("CAHSPER_ACCESS_TOKEN:", os.Getenv("CAHSPER_ACCESS_TOKEN"))
	},
}
