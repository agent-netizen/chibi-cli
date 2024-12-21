package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "chibi",
	Long: "Chibi for AniList - A lightweight anime & manga tracker CLI app powered by AniList.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(
		loginCmd,
		profileCmd,
		mediaSearchCmd,
		mediaListCmd,
		mediaUpdateCmd,
		mediaAddCmd,
	)
	rootCmd.Execute()
}
