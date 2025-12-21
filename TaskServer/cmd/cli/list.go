package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		req, _ := http.NewRequest("GET", baseURL+"/tasks", nil)
		req.Header.Set("X-API-Key", apiKey)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		b, _ := io.ReadAll(resp.Body)
		fmt.Println(string(b))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
