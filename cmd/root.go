package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var cli = &cobra.Command{
	Use: "dd",
	Short: "DefectDojo cli for CI/CD",
	Long: `This application is an application CLI for defectdojo`,
}

func init() {
}

func Execute() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}

