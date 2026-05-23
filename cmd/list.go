package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all subscriptions",
	Run: runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList(cmd *cobra.Command, args []string) {
	subs, err := jsonStorage.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(subs) == 0 {
		fmt.Println("No subscriptions found")
		return
	}

	fmt.Printf("%-20s | %8s | %-10s | %s\n", "Name", "Price", "Cycle", "Created")
	fmt.Println(strings.Repeat("-", 55))

	for _, sub := range subs {
		fmt.Println(sub)
	}
}
