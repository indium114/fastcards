package cmd

import (
	"github.com/spf13/cobra"
)

// archiveCmd represents the archive command
var archiveCmd = &cobra.Command{
	Use:   "archive <subcommand>",
	Short: "Manage the Deck Archive",
}

func init() {
	rootCmd.AddCommand(archiveCmd)
}
