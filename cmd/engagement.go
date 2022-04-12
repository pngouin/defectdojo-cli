package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/pngouin/defectdojo-cli/client"
	"github.com/pngouin/defectdojo-cli/config"
	"github.com/spf13/cobra"
)

var rootEngagement = &cobra.Command{
	Use:   "engagement",
	Short: "Engagement management",
	Long:  "Command to manage engagement",
}

var createEngagement = &cobra.Command{
	Use:   "create",
	Short: "create an engagement",
	Long:  "create an engagement for a product",
	Run:   create,
}

var closeEngagement = &cobra.Command{
	Use:   "close",
	Short: "close an engagement",
	Long:  "close an engagement for a product",
	Run:   close,
}

func init() {
	createEngagement.PersistentFlags().StringP("name", "n", "CI/CD", "engagement name")
	createEngagement.PersistentFlags().StringP("description", "d", "CI/CD engagement", "engagement description")
	createEngagement.PersistentFlags().StringP("engagement", "e", "CI/CD", "engagement name")
	createEngagement.PersistentFlags().StringP("version", "v", "", "version of the code repository")
	createEngagement.PersistentFlags().StringP("commit", "c", "", "commit hash of the code repository")
	createEngagement.PersistentFlags().StringP("branch", "b", "", "branch of the code repository")
	createEngagement.PersistentFlags().IntP("product", "p", -1, "product id")

	closeEngagement.PersistentFlags().StringP("engagement", "e", "", "engagement id")

	rootEngagement.AddCommand(createEngagement)
	rootEngagement.AddCommand(closeEngagement)
	cli.AddCommand(rootEngagement)
}

func create(cmd *cobra.Command, args []string) {
	timeStart := time.Now().Format("2006-01-02")
	timeEnd := time.Now().Add(time.Hour * 3).Format("2006-01-02")

	engagement := client.Engagement{
		ProductId:   getFlagInt(cmd, "product"),
		TargetStart: timeStart,
		TargetEnd:   timeEnd,
		Name:        getFlagS(cmd, "name"),
		Description: getFlagS(cmd, "description"),
		Version:     getFlagS(cmd, "version"),
		Status:      client.InProgress,
		Type:        client.CICD,
		CommitHash:  getFlagS(cmd, "commit"),
		BranchTag:   getFlagS(cmd, "branch"),
	}
	client := client.NewEngagementClient(config.Configuration)
	resp, err := client.Create(engagement)
	if err != nil {
		log.Fatalf("cannot create engagement: %s", err)
	}

	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
}

func close(cmd *cobra.Command, args []string) {
	client := client.NewEngagementClient(config.Configuration)
	engagement := getFlagS(cmd, "engagement")
	err := client.Close(engagement)
	if err != nil {
		log.Fatalf("cannot close engagement %s: %v", engagement, err)
	}
	fmt.Println("deleted")
}