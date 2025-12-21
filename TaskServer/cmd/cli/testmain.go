package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func testmain() {
	baseURL := flag.String("base", "http://localhost:8080", "API base URL")
	apiKey := flag.String("key", "dev-key", "API key")
	cmd := flag.String("cmd", "", "command: create | list | get")
	title := flag.String("title", "", "task title (for create)")
	id := flag.String("id", "", "task id (for get)")
	flag.Parse()

	client := &http.Client{}

	switch *cmd {
	case "create":
		if *title == "" {
			fmt.Println("title is required")
			os.Exit(1)
		}
		createTask(client, *baseURL, *apiKey, *title)

	case "list":
		listTasks(client, *baseURL, *apiKey)

	case "get":
		if *id == "" {
			fmt.Println("id is required")
			os.Exit(1)
		}
		getTask(client, *baseURL, *apiKey, *id)

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println(`
Usage:
  go run ./cmd/cli -cmd=create -title="hello"
  go run ./cmd/cli -cmd=list
  go run ./cmd/cli -cmd=get -id=<task-id>
`)
}

func createTask(c *http.Client, base, key, title string) {
	body, _ := json.Marshal(map[string]string{
		"title": title,
	})

	req, _ := http.NewRequest("POST", base+"/tasks", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", key)

	do(c, req)
}

func listTasks(c *http.Client, base, key string) {
	req, _ := http.NewRequest("GET", base+"/tasks", nil)
	req.Header.Set("X-API-Key", key)

	do(c, req)
}

func getTask(c *http.Client, base, key, id string) {
	req, _ := http.NewRequest("GET", base+"/tasks/"+id, nil)
	req.Header.Set("X-API-Key", key)

	do(c, req)
}

func do(c *http.Client, req *http.Request) {
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("request error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))
}
