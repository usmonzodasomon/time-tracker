package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/usmonzodasomon/time-tracker/internal/model"
)

type TaskRepoI interface {
	CreateTask(task model.Task) (int, error)
	GetTask(id int) (model.Task, error)
	StartTask(taskID int) error
	IsTaskStarted(taskID int) (bool, error)
	IsTaskStopped(taskID int) (bool, error)
	StopTask(id int) error
}

type TaskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) CreateTask(task model.Task) (int, error) {
	q := `INSERT INTO tasks (user_id, name, description)
	VALUES ($1, $2, $3) RETURNING id`
	if err := r.db.QueryRowx(q, task.UserID, task.Name, task.Description).
		Scan(&task.ID); err != nil {
		return 0, err
	}
	return task.ID, nil
}

func (r *TaskRepo) GetTask(id int) (model.Task, error) {
	q := `SELECT id, user_id, name, description FROM tasks WHERE id = $1`
	task := model.Task{}
	if err := r.db.Get(&task, q, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, model.ErrTaskNotFound
		}
		return model.Task{}, err
	}
	return task, nil
}

func (r *TaskRepo) StartTask(taskID int) error {
	q := `INSERT INTO time_entries (task_id, start_time) VALUES ($1, NOW())`
	_, err := r.db.Exec(q, taskID)
	return err
}

func (r *TaskRepo) IsTaskStarted(taskID int) (bool, error) {
	q := `SELECT id FROM time_entries WHERE task_id = $1 AND end_time IS NULL`
	var id int
	if err := r.db.Get(&id, q, taskID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *TaskRepo) IsTaskStopped(taskID int) (bool, error) {
	q := `SELECT id FROM time_entries WHERE task_id = $1 AND end_time IS NOT NULL`
	var id int
	if err := r.db.Get(&id, q, taskID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *TaskRepo) StopTask(taskID int) error {
	q := `UPDATE time_entries SET end_time = NOW() WHERE task_id = $1 AND end_time IS NULL`
	_, err := r.db.Exec(q, taskID)
	if err != nil {
		return err
	}
	return nil
}
