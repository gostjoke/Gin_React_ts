package main

import (
	"log"
	"net/http"

	"TaskServer/internal/httpapi"
	"TaskServer/internal/middleware"
	"TaskServer/internal/store"
	"TaskServer/internal/worker"
)

func main() {
	mem := store.NewMemoryStore()
	w := worker.New(mem, 100)

	mux := http.NewServeMux()
	api := httpapi.New(mem, w)

	mux.HandleFunc("GET /health", api.Health)
	mux.HandleFunc("POST /tasks", api.CreateTask)
	mux.HandleFunc("GET /tasks/{id}", api.GetTask)
	mux.HandleFunc("GET /tasks", api.ListTasks)

	// add handlers with middleware
	handler := middleware.Chain(
		mux,
		middleware.Logger(),
		middleware.APIKeyAuth("dev-key"),
	)

	// srv := &http.Server{
	// 	Addr:              ":8080",
	// 	Handler:           mux,
	// 	ReadHeaderTimeout: 5 * time.Second,
	// }

	srv := &http.Server{Addr: ":8080", Handler: handler /*...*/}

	log.Println("listening on :8080")
	log.Fatal(srv.ListenAndServe())
}
