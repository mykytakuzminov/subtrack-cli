package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mykytakuzminov/subtrack-cli/models"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a new subscription",
	Run: runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().String("name", "", "Subscription name")
	addCmd.Flags().Float64("price", 0, "Subscription price")
	addCmd.Flags().String("cycle", "", "Billing cycle (monthly/yearly)")
}

func runAdd(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	if name == "" {
		fmt.Println("Error: name is required")
		return
	}

	price, _ := cmd.Flags().GetFloat64("price")
	if price <= 0 {
		fmt.Println("Error: price must be positive")
		return
	}

	cycle, _ := cmd.Flags().GetString("cycle")
	if cycle != "monthly" && cycle != "yearly" {
		fmt.Println("Error: cycle must be 'monthly' or 'yearly'")
		return
	}

	s := models.NewSubscription(name, price, cycle)

	if err := jsonStorage.Save(s); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Subscription '%s' added successfully!\n", name)
}

