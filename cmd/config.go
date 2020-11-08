package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

		fmt.Print("UserName: ")
		scanner.Scan()
		userName := scanner.Text()
		if userName == "" {
			fmt.Print("UserName required.")
			os.Exit(0)
		}

		var err error
		var password = ""
		userName, password, err = utils.GetAccount(userName)
		if err != nil {
			if strings.Contains(fmt.Sprintln(err), "secret not found in keyring") {
				fmt.Printf("UserName %s does not exists. Continue to creating new account by %s.\n", userName, userName)
			} else {
				log.Fatal(err)
				os.Exit(1)
			}
		}
		utils.InteractInputHelper("Password", utils.Account, userName, password)

		idToken, err := utils.GetCredential(userName, utils.IDToken)
		if err != nil {
			if !strings.Contains(fmt.Sprintln(err), "secret not found in keyring") {
				log.Fatal(err)
				os.Exit(1)
			}
		}
		utils.InteractInputHelper("IdToken", utils.IDToken, userName, idToken)

		accessToken, err := utils.GetCredential(userName, utils.AccessToken)
		if err != nil {
			if !strings.Contains(fmt.Sprintln(err), "secret not found in keyring") {
				log.Fatal(err)
				os.Exit(1)
			}
		}
		utils.InteractInputHelper("AccessToken", utils.AccessToken, userName, accessToken)
	},
}
