package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/YoshinoriN/cahsper-cli/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(commentCommand)
	commentCommand.AddCommand(postCommentCommand)
}

var commentCommand = &cobra.Command{
	Use:   "comment",
	Short: "post your comment to cahsper",
	Long:  "post your comment to cahsper",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var postCommentCommand = &cobra.Command{
	Use:   "post",
	Short: "post your comment to cahsper",
	Long:  "post your comment to cahsper",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please input comment")
		}
		if len(args) > 101 {
			return errors.New("number of argument less than 100")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		cahsperConfigFilePath := utils.GetConfigFilePath()
		if !utils.Exists(cahsperConfigFilePath) {
			fmt.Println("config file not found. Please exec 'init'")
			return
		}
		cahsperConfig := utils.Read(cahsperConfigFilePath)

		accessToken, err := utils.GetCredential(cahsperConfig.Settings.Aws.Cognito.UserName, utils.AccessToken)
		if err != nil {
			if !strings.Contains(fmt.Sprintln(err), "secret not found in keyring") {
				log.Fatal(err)
			}
		}

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Do you want to post a comment? [y/n]: ")
		scanner.Scan()
		yn := strings.ToLower(scanner.Text())

		if yn != "y" && yn != "n" {
			fmt.Print("Please input 'y' or 'n'")
			os.Exit(0)
		}
		if yn != "y" {
			os.Exit(0)
		}

		u, err := url.Parse(cahsperConfig.Settings.ServerURL)
		u.Path = path.Join(u.Path, "users", cahsperConfig.Settings.Aws.Cognito.UserName, "comments")

		comment := []byte(fmt.Sprintf(`{"comment":"%s"}`, strings.Join(args, " ")))
		request, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(comment))
		if err != nil {
			log.Fatal(err)
		}

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

		client := new(http.Client)
		response, err := client.Do(request)

		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		fmt.Println("")
		fmt.Printf("Request URL: %s\n", response.Request.URL)
		fmt.Printf("Proto: %s\n", response.Proto)
		fmt.Printf("Date: %s\n", response.Header.Get("Date"))
		fmt.Printf("Server: %s\n", response.Header.Get("Server"))
		fmt.Printf("Request Method: %s\n", response.Request.Method)
		fmt.Printf("Status Code: %s\n", strconv.Itoa(response.StatusCode))
		fmt.Printf("Content-Length: %s\n", strconv.FormatInt(response.ContentLength, 10))
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	},
}
