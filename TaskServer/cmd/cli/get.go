package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Get a task by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		req, _ := http.NewRequest("GET", baseURL+"/tasks/"+id, nil)
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
	rootCmd.AddCommand(getCmd)
}
