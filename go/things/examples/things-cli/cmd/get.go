package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get and prints data for an instance of a thing",
	Long: `Get and prints data for an instance of a thing. Optionally, save to file the output.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := createConn()
		if err != nil { er(err) }
		log.Println(conn)

		if len(args) != 1 {
			er("invalid args count")
		}

		fmt.Println("get called")
	},
}

func init() {
	RootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&thingId, "id", "t", "", "")
}
