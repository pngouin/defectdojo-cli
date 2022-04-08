package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pngouin/defectdojo-cli/client"
	"github.com/pngouin/defectdojo-cli/config"
	"github.com/spf13/cobra"
)

var rootImportScan = &cobra.Command{
	Use: "scan",
	Short: "Import Scan management",
	Long: "Management API to Scan import",
}

var sendScan = &cobra.Command{
	Use: "send",
	Short: "send scan report",
	Long: "send scan report for an engagement",
	Run: send,
}

var listScan = &cobra.Command{
	Use: "list",
	Short: "list all scanner type",
	Long: "list all scanner type",
	Run: list,
}

func init() {
	sendScan.PersistentFlags().StringP("scan-type", "s", "", "name of the scanner use. dd scan list to get all the scanners")
	sendScan.PersistentFlags().StringP("file", "f", "", "path to the scanner output file")
	sendScan.PersistentFlags().IntP("engagement", "e", 0, "engagement id")

	rootImportScan.AddCommand(sendScan)
	rootImportScan.AddCommand(listScan)
	cli.AddCommand(rootImportScan)
}

func send(cmd *cobra.Command, args []string) {
	scanImport := client.ImportScan{
		EngagementId: getFlagInt(cmd, "engagement"),
		ScanType: getFlagS(cmd, "scan-type"),
		File: getFlagS(cmd, "file"),
	}

	client := client.NewImportScanClient(config.Configuration)
	res, err := client.Send(scanImport)
	if err != nil {
		log.Fatalf("cannot send scan: %v", err)
	}
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
}

func list(cmd *cobra.Command, args []string) {
	client := client.NewImportScanClient(config.Configuration)
	client.ListScan(false)
}