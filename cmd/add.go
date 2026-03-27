package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/indium114/fastcards/internal"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a card to a deck",
	Run: func(cmd *cobra.Command, args []string) {
		deck, err := internal.LoadDeck(args[0])
		if err != nil {
			fmt.Println("Deck not found")
			return
		}

		card := internal.Card{
			ID:    internal.NewID(),
			Front: args[1],
			Back:  args[2],
			State: 1,
		}

		deck.Cards = append(deck.Cards, card)

		if err := internal.SaveDeck(deck); err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println("Card added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
