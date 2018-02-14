package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search",
	Long: `Search`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := createConn()
		if err != nil { er(err) }

		if len(args) == 0 {
			er("invalid args count")
		}

		log.Println(searchFields, searchFilter, searchOpts, searchNS)
		q := things.NewStringQuery(searchFilter, searchOpts, searchFields)

		sr, err := conn.Search(searchNS, q)
		if err != nil { er(err) }

		log.Println(sr)

		fmt.Println("######### Found %s results ########", len(sr.Items))

		for i, v := range sr.Items {
			fmt.Println("  [%s] %s ", i, v.ThingId)
		}
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVar(&searchNS, "ns", "", "")
	searchCmd.Flags().StringVar(&searchFilter, "filter", "", "")
	searchCmd.Flags().StringVar(&searchOpts, "opts", "", "")
	searchCmd.Flags().StringVar(&searchFields, "fields", "", "")
}
