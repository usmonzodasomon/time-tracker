package external_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/usmonzodasomon/time-tracker/internal/model"
	"io"
	"net/http"
	"os"
)

type UserExternalInfoI interface {
	GetUser(passportSerie, passportNumber int) (model.User, error)
}

type UserExternalInfo struct {
	client *http.Client
}

func NewUserExternalInfo(client *http.Client) *UserExternalInfo {
	return &UserExternalInfo{client: client}
}

var URL = os.Getenv("EXTERNAL_API_URL")

func (u *UserExternalInfo) GetUser(passportSerie, passportNumber int) (model.User, error) {
	url := fmt.Sprintf("%s/info?passportSerie=%d&passportNumber=%d", URL, passportSerie, passportNumber)

	resp, err := u.client.Get(url)
	if err != nil {
		return model.User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.User{}, errors.New("error getting user info")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.User{}, err
	}

	user := struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic"`
		Address    string `json:"address"`
	}{}

	if err := json.Unmarshal(body, &user); err != nil {
		return model.User{}, err
	}

	return model.User{
		PassportSerie:  passportSerie,
		PassportNumber: passportNumber,
		Name:           user.Name,
		Surname:        user.Surname,
		Patronymic:     user.Patronymic,
		Address:        user.Address,
	}, nil
}
