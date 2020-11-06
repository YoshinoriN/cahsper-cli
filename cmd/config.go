package cmd

import (
	"bufio"
	"fmt"
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

		fmt.Print("Do you want to create a new credential? [y/n]: ")
		var yn = "y"
		for scanner.Scan() {
			yn = strings.ToLower(scanner.Text())
			break
		}

		if yn != "y" && yn != "n" {
			fmt.Print("Please input 'y' or 'n'")
			os.Exit(0)
		}

		var userName = ""
		var password = ""
		if yn == "y" {
			fmt.Println("Please input new UserName.")
		} else {
			fmt.Println("Please input exist UserName.")
		}
		fmt.Print("UserName: ")
		for scanner.Scan() {
			userName = scanner.Text()
			break
		}

		if userName == "" {
			fmt.Print("UserName required.")
			os.Exit(0)
		}

		var err error
		if yn == "n" && userName != "" {
			userName, password, err = utils.GetAccount(userName)
			if err != nil {
				fmt.Printf("UserName %s does not exists. Please input exists UserName or create new one.", userName)
				os.Exit(0)
			}
		}

		fmt.Printf("Password: %s", password)
		for scanner.Scan() {
			password = scanner.Text()
			break
		}
		err = utils.SetCredencial(userName, utils.Account, password)
		if err != nil {
			os.Exit(1)
		}

		fmt.Print("IdToken: ")
		for scanner.Scan() {
			idToken := scanner.Text()
			err = utils.SetCredencial(userName, utils.IDToken, idToken)
			if err != nil {
				os.Exit(1)
			}
			break
		}

		print("AccessToken: ")
		for scanner.Scan() {
			accessToken := scanner.Text()
			err = utils.SetCredencial(userName, utils.AccessToken, accessToken)
			if err != nil {
				os.Exit(1)
			}
			break
		}
	},
}
