package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func getFlagS(cmd *cobra.Command, flagName string) string {
	flagContent, err := cmd.Flags().GetString(flagName)
	if err != nil {
		log.Fatalf("cannot retrieve flag: %s", err)
	}
	return flagContent
}

func getFlagInt(cmd *cobra.Command, flagName string) int {
	flagContent, err := cmd.Flags().GetInt(flagName)
	if err != nil {
		log.Fatalf("cannot retrieve flag: %s", err)
	}
	return flagContent
}