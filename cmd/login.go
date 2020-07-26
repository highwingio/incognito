package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	auth "github.com/highwingio/incognito/internal/authentication"
	"github.com/highwingio/incognito/internal/storage"
	"github.com/highwingio/incognito/internal/types"
)

var loginUsername string
var loginPassword string
var loginClientID string
var loginUserPoolID string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Set up your Cognito credentials for generating tokens",
	Long:  "Store your Cognito credentials for generating tokens later on. Credentials are stored securely via your local Keyring.",
	RunE: func(cmd *cobra.Command, args []string) error {

		creds := &types.CognitoAuthenticationCredentials{
			Username:   loginUsername,
			Password:   loginPassword,
			UserPoolID: loginUserPoolID,
			ClientID:   loginClientID,
		}

		store, err := storage.NewKeyringStorage()
		if err != nil {
			return err
		}

		authenticator := auth.NewCognitoAuthenticator(creds)

		_, err = authenticator.SignIn()
		if err != nil {
			return err
		}

		fmt.Fprintln(os.Stderr, "Login successful!")

		err = store.StoreLoginCredentials(creds)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&loginUsername, "username", "u", "", "Cognito username to use for login")
	loginCmd.MarkFlagRequired("username")
	loginCmd.Flags().StringVarP(&loginPassword, "password", "p", "", "Password corresponding to the username")
	loginCmd.MarkFlagRequired("password")
	loginCmd.Flags().StringVarP(&loginClientID, "client", "c", "", "Cognito Client ID to authenticate with")
	loginCmd.MarkFlagRequired("client")
	loginCmd.Flags().StringVarP(&loginUserPoolID, "user-pool", "l", "", "Cognito User Pool ID to authenticate with")
	loginCmd.MarkFlagRequired("user-pool")
}
