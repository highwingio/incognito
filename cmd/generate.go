package cmd

import (
	"fmt"

	auth "github.com/highwingio/incognito/internal/authentication"
	"github.com/highwingio/incognito/internal/storage"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new Cognito token based on saved credentials",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := storage.NewKeyringStorage()
		if err != nil {
			return err
		}

		creds, err := store.RetrieveLoginCredentials()
		if err != nil {
			return err
		}

		authenticator := auth.NewCognitoAuthenticator(creds)

		result, err := authenticator.SignIn()
		if err != nil {
			return err
		}

		fmt.Print(*result.AccessToken)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
