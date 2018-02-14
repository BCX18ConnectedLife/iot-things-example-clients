package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things"
)

var subCmd = &cobra.Command{
	Use: "sub [scope]",
	Aliases: []string{"subscribe "},
	Short: "Subscribes to Things",
	Long: `Subscribes and prints out to console incoming messages`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := createConn()
		if err != nil { er(err) }

		obsEvents := make(chan *things.WSMessage)
		conn.ObserveEvents(obsEvents)
		fmt.Println("Started Observing..")

		for {
			select {
			case obsMsg, _ := <-obsEvents:
				fmt.Println(">> ", obsMsg.Topic)
			}
		}
	},
}

var scope string

func init() {
	RootCmd.AddCommand(subCmd)

	subCmd.Flags().StringVarP(&scope, "scope", "s", "all", "The scope of subscription, what to listen to. Defaults to 'all'")
}
