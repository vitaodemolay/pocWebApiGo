package service

import (
	"errors"
	"web-api-gin/model"
)

var errLoginNotFound = errors.New("login not found")

var dbLogins = []model.Login{}

func GetLoginByToken(token string) (err error, login model.Login) {
	for _, l := range dbLogins {
		if l.UserId == token {
			login = l
			err = nil
			return
		}
	}
	err = errLoginNotFound
	return
}

func AddLogin(user model.User) (login model.Login) {
	removeLoginByUserName(user.UserName)
	login = model.CreateLogin(user)
	dbLogins = append(dbLogins, login)
	return
}

func RemoveLoginByToken(token string) error {
	for idx, l := range dbLogins {
		if l.UserId == token {
			dbLogins = append(dbLogins[:idx], dbLogins[idx+1:]...)
			return nil
		}
	}
	return errLoginNotFound
}

func RemoveLoginExpired(limitInSeconds float64) {
	for idx, l := range dbLogins {
		if l.IsExpired(limitInSeconds) {
			dbLogins = append(dbLogins[:idx], dbLogins[idx+1:]...)
		}
	}
}

func removeLoginByUserName(userName string) {
	for idx, l := range dbLogins {
		if l.User.UserName == userName {
			dbLogins = append(dbLogins[:idx], dbLogins[idx+1:]...)
			break
		}
	}
}
