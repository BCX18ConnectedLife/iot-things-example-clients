package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"log"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new thing",
	Long: `Add a new thing`,
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
		// call add


		fmt.Println("add called")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&thingId, "id", "t", "", "")
	addCmd.Flags().StringVarP(&content, "content", "c", "", "")
	addCmd.Flags().StringVarP(&file, "file", "f", "", "")
}
