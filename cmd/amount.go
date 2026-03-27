package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/indium114/fastcards/internal"
)

// amountCmd represents the amount command
var amountCmd = &cobra.Command{
	Use:   "amount [deck]",
	Short: "Print the amount of cards currently due",
	RunE: func(cmd *cobra.Command, args []string) error {
		var decks []*internal.Deck

		// If deck specified
		if len(args) == 1 {
			d, err := internal.LoadDeck(args[0])
			if err != nil {
				return fmt.Errorf("Deck '%s' not found", args[0])
			}
			decks = append(decks, &d)
		} else {
			// Load all decks
			names, err := internal.ListDeckNames()
			if err != nil {
				return err
			}

			for _, name := range names {
				d, err := internal.LoadDeck(name)
				if err == nil {
					decks = append(decks, &d)
				}
			}
		}

		total := 0

		for _, d := range decks {
			for _, c := range d.Cards {
				if internal.IsDue(c) {
					total++
				}
			}
		}

		fmt.Println("󰘸 Flashcards due:", total)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(amountCmd)
}
