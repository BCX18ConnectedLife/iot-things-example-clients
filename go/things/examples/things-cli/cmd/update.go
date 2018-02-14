package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a thing instance",
	Long: `Updates a thing instance`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := createConn()
		if err != nil { er(err) }
		log.Println(conn)

		if len(args) > 2 || len(args) < 2{
			er("invalid args count")
		}

		// if file
			// load content from file

		// validate json content
		// call update

		fmt.Println("update called")
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&thingId, "id", "t", "", "")
	updateCmd.Flags().StringVarP(&content, "content", "c", "", "")
	updateCmd.Flags().StringVarP(&file, "file", "f", "", "")
}
