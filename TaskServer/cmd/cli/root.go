package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	baseURL string
	apiKey  string
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task CLI for TaskServer",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&baseURL,
		"base",
		"http://localhost:8080",
		"API base URL",
	)

	rootCmd.PersistentFlags().StringVar(
		&apiKey,
		"key",
		"dev-key",
		"API key",
	)
}
