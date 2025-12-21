package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"TaskServer/internal/model"
	"TaskServer/internal/store"
)

type Store interface {
	Create(model.Task) (model.Task, error)
	Get(id string) (model.Task, error)
	List() ([]model.Task, error)
	Update(model.Task) (model.Task, error)
}

type API struct{ store Store }

func New(s Store) *API { return &API{store: s} }

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func (a *API) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]string{"ok": "true"})
}

type createTaskReq struct {
	Title string `json:"title"`
}

func (a *API) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req createTaskReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" {
		writeJSON(w, 400, map[string]string{"error": "invalid json or empty title"})
		return
	}

	t := model.Task{
		ID:        newID(),
		Title:     req.Title,
		Status:    model.StatusQueued,
		CreatedAt: time.Now(),
	}

	created, err := a.store.Create(t)
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": "create failed"})
		return
	}
	writeJSON(w, 201, created)
}

func (a *API) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	t, err := a.store.Get(id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, 500, map[string]string{"error": "get failed"})
		return
	}
	writeJSON(w, 200, t)
}

func (a *API) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := a.store.List()
	if err != nil {
		writeJSON(w, 500, map[string]string{"error": "list failed"})
		return
	}
	writeJSON(w, 200, tasks)
}
