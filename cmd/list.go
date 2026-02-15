package cmd

import (
	"fmt"

	"github.com/stikypiston/fastcards/internal"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all decks and number of cards",
	RunE: func(cmd *cobra.Command, args []string) error {

		names, err := internal.ListDeckNames()
		if err != nil {
			return err
		}

		if len(names) == 0 {
			fmt.Println("No decks found.")
			return nil
		}

		for _, name := range names {
			deck, err := internal.LoadDeck(name)
			if err != nil {
				continue
			}
			total := len(deck.Cards)
			due := 0
			for _, c := range deck.Cards {
				if internal.IsDue(c) {
					due++
				}
			}
			fmt.Printf("%s: %d cards (%d due)\n", name, total, due)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
