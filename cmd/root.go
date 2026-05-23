package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mykytakuzminov/subtrack-cli/storage"
)

var jsonStorage = storage.NewJSONStorage("subscriptions.json")

var rootCmd = &cobra.Command{
	Use:   "subtrack",
	Short: "Track your subscriptions",
}

func Execute() {
	rootCmd.Execute()
}
