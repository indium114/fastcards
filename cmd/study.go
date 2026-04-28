package cmd

import (
	"fmt"
	"github.com/indium114/fastcards/internal"
	"github.com/indium114/fastcards/ui"

	"github.com/spf13/cobra"

	tea "github.com/charmbracelet/bubbletea"
)

// studyCmd represents the study command
var studyCmd = &cobra.Command{
	Use:   "study [deck]",
	Short: "Study due cards",
	Args:  cobra.MaximumNArgs(1),
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

		// Collect all due cards
		var due []internal.DueRef

		for _, d := range decks {
			for i, c := range d.Cards {
				if internal.IsDue(c) {
					due = append(due, internal.DueRef{
						Deck: d,
						Idx:  i,
					})
				}
			}
		}

		if len(due) == 0 {
			fmt.Println("No cards due.")
			return nil
		}

		// Start Bubble Tea
		p := tea.NewProgram(ui.NewStudyModelFromRefs(due))
		if _, err := p.Run(); err != nil {
			return err
		}

		// Save all decks after study
		for _, d := range decks {
			internal.SaveDeck(*d)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(studyCmd)
}
