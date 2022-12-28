package model

import (
	"time"

	"github.com/google/uuid"
)

type Login struct {
	UserId   string `json:"userId" description:"Authentication Token"`
	User     User   `json:"user"`
	Register time.Time
}

func CreateLogin(user User) (login Login) {
	login.User = user
	login.UserId = uuid.New().String()
	login.Register = time.Now()
	return
}

func (login Login) IsExpired(limitInSeconds float64) bool {
	return time.Now().Sub(login.Register).Seconds() > limitInSeconds
}
