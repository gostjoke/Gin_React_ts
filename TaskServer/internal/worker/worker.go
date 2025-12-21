package worker

import (
	"context"
	"log"
	"time"

	"TaskServer/internal/model"
)

type Store interface {
	Get(id string) (model.Task, error)
	Update(model.Task) (model.Task, error)
}

type Worker struct {
	store Store
	q     chan string
}

func New(store Store, buf int) *Worker {
	return &Worker{store: store, q: make(chan string, buf)}
}

func (w *Worker) Enqueue(taskID string) {
	select {
	case w.q <- taskID:
	default:
		// queue full: 你可以改成返回 error 或 drop / metrics
		log.Println("queue full, dropping:", taskID)
	}
}

func (w *Worker) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case id := <-w.q:
			w.process(id)
		}
	}
}

func (w *Worker) process(id string) {
	t, err := w.store.Get(id)
	if err != nil {
		log.Println("get task failed:", err)
		return
	}
	t.Status = model.StatusRunning
	_, _ = w.store.Update(t)

	// 模擬耗時工作
	time.Sleep(800 * time.Millisecond)

	t.Status = model.StatusDone
	_, _ = w.store.Update(t)
}
