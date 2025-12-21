package model

import "time"

type TaskStatus string

const (
	StatusQueued  TaskStatus = "queued"
	StatusRunning TaskStatus = "running"
	StatusDone    TaskStatus = "done"
	StatusFailed  TaskStatus = "failed"
)

type Task struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	Title     string     `json:"title"`
	Status    TaskStatus `json:"status"`
}
