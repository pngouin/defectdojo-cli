package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pngouin/defectdojo-cli/client"
	"github.com/pngouin/defectdojo-cli/config"
	"github.com/spf13/cobra"
)

var rootLanguageScan = &cobra.Command{
	Use: "languages",
	Short: "Import languages report",
	Long: "Import languages report from cloc",
}

var sendLanguages = &cobra.Command{
	Use: "send",
	Short: "send languages report",
	Long: "Send languages report from cloc in format json",
	Run: sendLanguagesF,
}

func init() {
	sendLanguages.PersistentFlags().StringP("file", "f", "", "path to cloc output file")
	sendLanguages.PersistentFlags().IntP("product", "p", -1, "product id")

	rootLanguageScan.AddCommand(sendLanguages)
	cli.AddCommand(rootLanguageScan)
}

func sendLanguagesF(cmd *cobra.Command, args []string) {
	languagesReport := client.ImportLanguages{
		Product: getFlagInt(cmd, "product"),
		File: getFlagS(cmd, "file"),
	}

	client := client.NewImportLanguagesClient(config.Configuration)
	res, err := client.Send(languagesReport)
	if err != nil {
		log.Fatalf("cannot send languages report: %v", err)
	}
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
}