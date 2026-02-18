package cmd

import (
	"fmt"

	"github.com/stikypiston/fastcards/internal"

	"github.com/spf13/cobra"
)

var archivelistCmd = &cobra.Command{
	Use:   "list",
	Short: "List all decks and number of cards",
	RunE: func(cmd *cobra.Command, args []string) error {

		names, err := internal.ListArchivedDeckNames()
		if err != nil {
			return err
		}

		if len(names) == 0 {
			fmt.Println("No decks found.")
			return nil
		}

		for _, name := range names {
			fmt.Println(name)
		}

		return nil
	},
}

func init() {
	archiveCmd.AddCommand(archivelistCmd)
}
