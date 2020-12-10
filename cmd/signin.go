package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/YoshinoriN/cahsper-cli/utils"
	cognitosrp "github.com/alexrudd/cognito-srp/v2"
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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires UserName")
		}
		if len(args) > 1 {
			return errors.New("number of argument must be one")
		}
		return nil
	},
	// TODO: refactor
	Run: func(cmd *cobra.Command, args []string) {
		userName, password, err := utils.GetAccount(args[0])
		if err != nil {
			if strings.Contains(fmt.Sprintln(err), "secret not found in keyring") {
				fmt.Printf("UserName %s does not exists.\n", userName)
				os.Exit(0)
			} else {
				log.Fatal(err)
				os.Exit(1)
			}
		}

		cahsperConfigFilePath := utils.GetConfigFilePath()
		if !utils.Exists(cahsperConfigFilePath) {
			fmt.Println("config file not found. Please exec 'init'")
			return
		}
		cahsperConfig := utils.Read(cahsperConfigFilePath)

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

		resp, err := svc.InitiateAuth(context.Background(), &cip.InitiateAuthInput{
			AuthFlow:       types.AuthFlowTypeUserSrpAuth,
			ClientId:       aws.String(csrp.GetClientId()),
			AuthParameters: csrp.GetAuthParams(),
		})
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		if resp.ChallengeName == types.ChallengeNameTypePasswordVerifier {
			challengeResponses, _ := csrp.PasswordVerifierChallenge(resp.ChallengeParameters, time.Now())

			resp, err := svc.RespondToAuthChallenge(context.Background(), &cip.RespondToAuthChallengeInput{
				ChallengeName:      types.ChallengeNameTypePasswordVerifier,
				ChallengeResponses: challengeResponses,
				ClientId:           aws.String(csrp.GetClientId()),
			})
			if err != nil {
				log.Fatal(err)
				panic(err)
			}

			// TODO: write to keyring
			fmt.Printf("Access Token: %s\n", *resp.AuthenticationResult.AccessToken)
			fmt.Printf("ID Token: %s\n", *resp.AuthenticationResult.IdToken)
			fmt.Printf("Refresh Token: %s\n", *resp.AuthenticationResult.RefreshToken)
		}
	},
}
