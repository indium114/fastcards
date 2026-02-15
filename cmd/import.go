package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/stikypiston/fastcards/internal"
)

var importCmd = &cobra.Command{
	Use:   "import [tsv_file]",
	Short: "Import flashcards from a TSV file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]

		f, err := os.Open(path)
		if err != nil {
			return fmt.Errorf(" Failed to open file: %w", err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		lineNo := 0

		for scanner.Scan() {
			lineNo++
			line := scanner.Text()
			if strings.TrimSpace(line) == "" {
				continue
			}

			parts := strings.Split(line, "\t")
			if len(parts) < 3 {
				fmt.Printf("󰒭 Skipping line %d: not enough columns\n", lineNo)
				continue
			}

			deckName := strings.TrimSpace(parts[0])
			front := strings.TrimSpace(parts[1])
			back := strings.TrimSpace(parts[2])

			if deckName == "" || front == "" || back == "" {
				fmt.Printf("󰒭 Skipping line %d: empty deck/front/back\n", lineNo)
				continue
			}

			internal.CreateDeck(deckName)
			deck, err := internal.LoadDeck(deckName)
			if err != nil {
				fmt.Println("Error:", err)
			}

			card := internal.Card{
				ID:           uuid.NewString(),
				Front:        front,
				Back:         back,
				State:        1,
				LastReviewed: nil,
			}

			deck.Cards = append(deck.Cards, card)

			if err := internal.SaveDeck(deck); err != nil {
				return fmt.Errorf("failed to save deck '%s': %w", deckName, err)
			}
		}

		if scanner.Err() != nil {
			return fmt.Errorf("error reading file: %w", scanner.Err())
		}

		fmt.Println("Import complete!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
