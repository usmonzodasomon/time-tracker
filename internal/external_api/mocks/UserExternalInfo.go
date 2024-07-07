package mocks

import (
	"errors"
	"github.com/usmonzodasomon/time-tracker/internal/model"
)

type UserExternalInfo struct {
	data []model.User
}

func NewUserExternalInfo() *UserExternalInfo {
	return &UserExternalInfo{
		data: []model.User{
			{
				PassportSerie:  1234,
				PassportNumber: 5678,
				Name:           "Петр",
				Surname:        "Петров",
				Patronymic:     "Петрович",
				Address:        "ул. Петрова, д. 1",
			},
			{
				PassportSerie:  4321,
				PassportNumber: 8765,
				Name:           "Иван",
				Surname:        "Иванов",
				Patronymic:     "Иванович",
				Address:        "ул. Иванова, д. 2",
			},
			{
				PassportSerie:  1111,
				PassportNumber: 2222,
				Name:           "Сидор",
				Surname:        "Сидоров",
				Patronymic:     "Сидорович",
				Address:        "ул. Сидорова, д. 3",
			},
			{
				PassportSerie:  3333,
				PassportNumber: 4444,
				Name:           "Александр",
				Surname:        "Александров",
				Patronymic:     "Александрович",
				Address:        "ул. Александрова, д. 4",
			},
			{
				PassportSerie:  5555,
				PassportNumber: 6666,
				Name:           "Алексей",
				Surname:        "Алексеев",
				Patronymic:     "Алексеевич",
				Address:        "ул. Алексеева, д. 5",
			},
			{
				PassportSerie:  7777,
				PassportNumber: 8888,
				Name:           "Андрей",
				Surname:        "Андреев",
				Patronymic:     "Андреевич",
				Address:        "ул. Андреева, д. 6",
			},
		},
	}
}

func (u *UserExternalInfo) GetUser(passportSerie, passportNumber int) (model.User, error) {
	for _, user := range u.data {
		if user.PassportSerie == passportSerie && user.PassportNumber == passportNumber {
			return user, nil
		}
	}
	return model.User{}, errors.New("user not found")
}
