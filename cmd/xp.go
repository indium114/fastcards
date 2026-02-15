package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stikypiston/fastcards/internal"
)

// xpCmd represents the xp command
var xpCmd = &cobra.Command{
	Use:   "xp",
	Short: "Show XP amount",
	RunE: func(cmd *cobra.Command, args []string) error {
		xp, err := internal.LoadXP()
		if err != nil {
			return err
		}
		fmt.Printf(" XP: %d\n", xp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(xpCmd)
}
