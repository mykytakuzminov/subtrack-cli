package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a subscription",
	Run:   runDelete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func runDelete(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: subscription ID is required")
		return
	}

	id := args[0]

	if err := jsonStorage.Delete(id); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Subscription deleted successfully!\n")
}
