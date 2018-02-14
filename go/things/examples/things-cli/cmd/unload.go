package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// unloadCmd represents the unload command
var unloadCmd = &cobra.Command{
	Use:   "unload",
	Short: "Given a JSON file of an array of things, delete each instance from Things",
	Long: `Given a JSON file of an array of things, delete each instance from Things`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := createConn()
		if err != nil { er(err) }
		log.Println(conn)

		if len(args) != 1 {
			er("invalid args count")
		}

		fmt.Println("unload called")
	},
}

func init() {
	RootCmd.AddCommand(unloadCmd)

	unloadCmd.Flags().StringVarP(&file, "file", "f", "", "")
}
