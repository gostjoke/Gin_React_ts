package main

import (
	"log"
	"net/http"
	"time"

	"TaskServer/internal/httpapi"
	"TaskServer/internal/store"
)

func main() {
	mem := store.NewMemoryStore()

	mux := http.NewServeMux()
	api := httpapi.New(mem)

	mux.HandleFunc("GET /health", api.Health)
	mux.HandleFunc("POST /tasks", api.CreateTask)
	mux.HandleFunc("GET /tasks/{id}", api.GetTask)
	mux.HandleFunc("GET /tasks", api.ListTasks)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("listening on :8080")
	log.Fatal(srv.ListenAndServe())
}
