package cmd

import (
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Stores login credentials",
	Long: `Stores login credentials for subsequent use with Things`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := &Config{}

		cfg.Username = username
		cfg.Password = password
		cfg.ApiToken = token
		cfg.Proxy = proxy
		cfg.Endpoint = endpoint

		err := createConfigFile(cfg)

		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&username, "user", "u", "", "")
	loginCmd.Flags().StringVarP(&password, "pass", "p", "", "")
	loginCmd.Flags().StringVarP(&token, "token", "t", "", "")
	loginCmd.Flags().StringVarP(&proxy, "proxy", "x", "", "")
	loginCmd.Flags().StringVarP(&endpoint, "ep", "e", "wss://things.s-apps.de1.bosch-iot-cloud.com", "")
}
