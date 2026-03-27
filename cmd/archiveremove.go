package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/indium114/fastcards/internal"
	"log"
	"os"
	"path/filepath"
)

var archiveremoveCmd = &cobra.Command{
	Use:   "remove <file>",
	Short: "Remove a deck from the Deck Archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deck := args[0] + ".json"

		newpath := filepath.Join(internal.DecksDir(), deck)
		oldpath := filepath.Join(internal.ArchiveDir(), deck)

		err := os.Rename(oldpath, newpath)
		if err != nil {
			log.Fatal("Error:", err)
		}

		fmt.Println("Unarchived deck:", deck)
	},
}

func init() {
	archiveCmd.AddCommand(archiveremoveCmd)
}
