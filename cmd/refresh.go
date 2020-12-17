package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/YoshinoriN/cahsper-cli/utils"
	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go-v2/aws"
	v1Aws "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cipv1 "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func init() {
	rootCommand.AddCommand(refreshTokenCommand)
}

var refreshTokenCommand = &cobra.Command{
	Use:   "refresh",
	Short: "refresh token of aws cognito",
	Long:  "refresh token of aws cognito",
	Args:  cobra.NoArgs,
	// TODO: refactor
	Run: func(cmd *cobra.Command, args []string) {
		cahsperConfigFilePath := utils.GetConfigFilePath()
		if !utils.Exists(cahsperConfigFilePath) {
			fmt.Println("config file not found. Please exec 'init'")
			return
		}
		cahsperConfig := utils.Read(cahsperConfigFilePath)

		// TODO: username existing validation
		refreshToken, err := utils.GetCredential(cahsperConfig.Settings.Aws.Cognito.UserName, utils.RefreshToken)
		if err != nil {
			if !strings.Contains(fmt.Sprintln(err), "secret not found in keyring") {
				log.Fatal(err)
			}
		}

		authInputParams := &cipv1.InitiateAuthInput{
			ClientId: aws.String(cahsperConfig.Settings.Aws.Cognito.AppClientID),
			AuthFlow: aws.String("REFRESH_TOKEN_AUTH"),
			AuthParameters: map[string]*string{
				"REFRESH_TOKEN": aws.String(refreshToken),
			},
		}

		mySession := session.Must(session.NewSession())
		svc := cipv1.New(mySession, v1Aws.NewConfig().WithRegion(cahsperConfig.Settings.Aws.Region))

		request, response := svc.InitiateAuthRequest(authInputParams)
		err = request.Send()
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		utils.SetCredential(cahsperConfig.Settings.Aws.Cognito.UserName, utils.AccessToken, *response.AuthenticationResult.AccessToken)
		utils.SetCredential(cahsperConfig.Settings.Aws.Cognito.UserName, utils.IDToken, *response.AuthenticationResult.IdToken)
	},
}
