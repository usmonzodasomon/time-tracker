package model

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID             int    `db:"id"`
	PassportSerie  int    `db:"passport_serie"`
	PassportNumber int    `db:"passport_number"`
	Name           string `db:"name"`
	Surname        string `db:"surname"`
	Patronymic     string `db:"patronymic"`
	Address        string `db:"address"`
}

type UserRequestBody struct {
	PassportNumber string `json:"passportNumber"`
}

type UserTaskTimeSpent struct {
	TaskID  int
	Hours   int
	Minutes int
}

type UserUpdateRequestBody struct {
	Name       *string `json:"name"`
	Surname    *string `json:"surname"`
	Patronymic *string `json:"patronymic"`
	Address    *string `json:"address"`
}

type UserFilter struct {
	ID             *int    `form:"id"`
	PassportSerie  *int    `form:"passport_serie"`
	PassportNumber *int    `form:"passport_number"`
	Name           *string `form:"name"`
	Surname        *string `form:"surname"`
	Patronymic     *string `form:"patronymic"`
	Address        *string `form:"address"`

	Page    int `form:"page" default:"1"`
	PerPage int `form:"per_page" default:"10"`
}
