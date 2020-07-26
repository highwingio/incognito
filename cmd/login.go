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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
