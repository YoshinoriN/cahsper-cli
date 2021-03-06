package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/yoshinorin/cahsper-cli/utils"
	cognitosrp "github.com/alexrudd/cognito-srp/v2"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func init() {
	rootCommand.AddCommand(signInCommand)
}

var signInCommand = &cobra.Command{
	Use:   "signin",
	Short: "signin aws cognito",
	Long:  "signin aws cognito",
	Args:  cobra.NoArgs,
	// TODO: refactor
	Run: func(cmd *cobra.Command, args []string) {
		cahsperConfigFilePath := utils.GetConfigFilePath()
		if !utils.Exists(cahsperConfigFilePath) {
			fmt.Println("config file not found. Please exec 'init'")
			return
		}
		cahsperConfig := utils.Read(cahsperConfigFilePath)

		userName, password, err := utils.GetAccount(cahsperConfig.Settings.Aws.Cognito.UserName)
		if err != nil {
			if strings.Contains(fmt.Sprintln(err), "secret not found in keyring") {
				fmt.Printf("UserName %s does not exists.\n", userName)
				os.Exit(0)
			} else {
				log.Fatal(err)
			}
		}

		csrp, _ := cognitosrp.NewCognitoSRP(
			userName,
			password,
			cahsperConfig.Settings.Aws.Cognito.UserPoolID,
			cahsperConfig.Settings.Aws.Cognito.AppClientID,
			nil,
		)

		cfg, _ := awsConfig.LoadDefaultConfig(
			awsConfig.WithRegion(cahsperConfig.Settings.Aws.Region),
			awsConfig.WithCredentialsProvider(aws.AnonymousCredentials{}),
		)
		svc := cip.NewFromConfig(cfg)

		response, err := svc.InitiateAuth(context.Background(), &cip.InitiateAuthInput{
			AuthFlow:       types.AuthFlowTypeUserSrpAuth,
			ClientId:       aws.String(csrp.GetClientId()),
			AuthParameters: csrp.GetAuthParams(),
		})
		if err != nil {
			log.Fatal(err)
		}

		if response.ChallengeName == types.ChallengeNameTypePasswordVerifier {
			challengeResponses, _ := csrp.PasswordVerifierChallenge(response.ChallengeParameters, time.Now())

			response, err := svc.RespondToAuthChallenge(context.Background(), &cip.RespondToAuthChallengeInput{
				ChallengeName:      types.ChallengeNameTypePasswordVerifier,
				ChallengeResponses: challengeResponses,
				ClientId:           aws.String(csrp.GetClientId()),
			})
			if err != nil {
				log.Fatal(err)
			}

			err = utils.SetCredential(userName, utils.AccessToken, *response.AuthenticationResult.AccessToken)
			if err != nil {
				log.Fatal(err)
			}
			err = utils.SetCredential(userName, utils.IDToken, *response.AuthenticationResult.IdToken)
			if err != nil {
				log.Fatal(err)
			}
			err = utils.SetCredential(userName, utils.RefreshToken, *response.AuthenticationResult.RefreshToken)
			if err != nil {
				log.Fatal(err)
			}
		}
		color.Set(color.FgHiGreen)
		fmt.Println("[succeeded]: successfully sign in and get a new auth token.")
		color.Unset()
	},
}
