package service

import (
	"fmt"
	"github.com/usmonzodasomon/time-tracker/internal/model"
	"github.com/usmonzodasomon/time-tracker/internal/repository"
	"time"
)

type UserServiceI interface {
	GetAllUsers(filter model.UserFilter) ([]model.User, error)
	GetUserTimeSpent(userID int, startPeriod, endPeriod time.Time) ([]model.UserTaskTimeSpent, error)

	GetUser(id int) (model.User, error)
	CreateUser(user model.User) (int, error)
	DeleteUser(id int) error
	UpdateUser(user model.User) error
}

type UserService struct {
	repo repository.UserRepoI
}

func NewUserService(repo repository.UserRepoI) *UserService {
	return &UserService{repo: repo}
}
func (s *UserService) GetAllUsers(filter model.UserFilter) ([]model.User, error) {
	return s.repo.GetAllUsers(filter)
}
func (s *UserService) GetUserTimeSpent(userID int, startPeriod, endPeriod time.Time) ([]model.UserTaskTimeSpent, error) {
	_, err := s.GetUser(userID)
	if err != nil {
		return nil, err
	}
	timeSpentMinutes, err := s.repo.GetUserTimeSpent(userID, startPeriod, endPeriod)
	if err != nil {
		return nil, fmt.Errorf("error getting user time spent: %w", err)
	}

	userTaskTimeSpent := make([]model.UserTaskTimeSpent, 0, len(timeSpentMinutes))
	for _, v := range timeSpentMinutes {
		hours := int(v.TotalMinutes) / 60
		minutes := int(v.TotalMinutes) % 60
		userTaskTimeSpent = append(userTaskTimeSpent, model.UserTaskTimeSpent{
			TaskID:  v.TaskID,
			Hours:   hours,
			Minutes: minutes,
		})
	}

	return userTaskTimeSpent, nil
}

func (s *UserService) GetUser(id int) (model.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserService) CreateUser(user model.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUser(user model.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	_, err := s.GetUser(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteUser(id)
}
