package cmd

import (
	"github.com/mykytakuzminov/subtrack-cli/storage"
	"github.com/spf13/cobra"
)

var jsonStorage = storage.NewJSONStorage("subscriptions.json")

var rootCmd = &cobra.Command{
	Use:   "subtrack",
	Short: "Track your subscriptions",
}

func Execute() {
	rootCmd.Execute()
}
