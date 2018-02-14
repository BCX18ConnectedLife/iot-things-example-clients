package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Reads from a JSON file all things and attempts to create an instance for each",
	Long: `Reads from a JSON file all things and attempts to create an instance for each`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := createConn()
		if err != nil { er(err) }
		log.Println(conn)

		if len(args) != 1 {
			er("invalid args count")
		}

		ts, err := loadThingsJsonFile(file)
		if err != nil { er(err) }
		fmt.Println(len(ts))

		for _, v := range ts {
			conn.Add(v)
		}
		fmt.Println("load called")
	},
}

func init() {
	RootCmd.AddCommand(loadCmd)

	loadCmd.Flags().StringVarP(&file, "file", "f", "", "")
}
