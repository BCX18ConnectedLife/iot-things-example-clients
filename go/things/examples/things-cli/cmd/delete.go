package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a thing instance",
	Long: `Deletes a thing instance`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := createConn()
		if err != nil { er(err) }
		log.Println(conn)

		if len(args) != 1 {
			er("invalid args count")
		}

		// call delete

		fmt.Println("delete called")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&thingId, "id", "t", "", "")
}
