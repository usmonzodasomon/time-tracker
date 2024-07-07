package service

import (
	"github.com/usmonzodasomon/time-tracker/internal/model"
	"github.com/usmonzodasomon/time-tracker/internal/repository"
)

type TaskServiceI interface {
	CreateTask(task model.Task) (int, error)
	GetTask(id int) (model.Task, error)
	StartTask(taskID int) error
	IsTaskStarted(taskID int) (bool, error)
	IsTaskStopped(taskID int) (bool, error)
	StopTask(id int) error
}

type TaskService struct {
	repo repository.TaskRepoI
}

func NewTaskService(repo repository.TaskRepoI) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task model.Task) (int, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetTask(id int) (model.Task, error) {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (s *TaskService) StartTask(taskID int) error {
	started, err := s.IsTaskStarted(taskID)
	if err != nil {
		return err
	}
	if started {
		return model.ErrTaskAlreadyStarted
	}
	return s.repo.StartTask(taskID)
}

func (s *TaskService) IsTaskStarted(taskID int) (bool, error) {
	return s.repo.IsTaskStarted(taskID)
}

func (s *TaskService) IsTaskStopped(taskID int) (bool, error) {
	return s.repo.IsTaskStopped(taskID)

}

func (s *TaskService) StopTask(id int) error {
	stopped, err := s.IsTaskStopped(id)
	if err != nil {
		return err
	}
	if stopped {
		return model.ErrTaskAlreadyStopped
	}

	return s.repo.StopTask(id)
}
