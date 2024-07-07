package model

import "errors"

var (
	ErrTaskNotFound       = errors.New("task not found")
	ErrTaskAlreadyStarted = errors.New("task already started")
	ErrTaskAlreadyStopped = errors.New("task already stopped")
)

type Task struct {
	ID          int    `db:"id"`
	UserID      int    `db:"user_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type TaskTimeSpent struct {
	TaskID       int     `db:"task_id"`
	TotalMinutes float64 `db:"total_minutes"`
}

type TaskRequestBody struct {
	UserID      int    `json:"user_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
