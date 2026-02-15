package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stikypiston/fastcards/internal"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new deck",
	Run: func(cmd *cobra.Command, args []string) {
		deck := internal.Deck{
			ID:    internal.NewID(),
			Name:  args[0],
			Cards: []internal.Card{},
		}

		if err := internal.SaveDeck(deck); err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Deck created:", deck.Name)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
