package service

import (
	"errors"
	"web-api-gin/model"
)

var errUserNotFound = errors.New("user not found")

var dbUsers = []model.User{
	{UserName: "admin", Password: "Passw0rd"},
}

func GetAllUsers() []model.User {
	return dbUsers
}

func AddNewUser(newUser model.User) {
	dbUsers = append(dbUsers, newUser)
}

func GetUserByUserName(userName string) (err error, user model.User) {
	for _, u := range dbUsers {
		if u.UserName == userName {
			user = u
			err = nil
			return
		}
	}
	err = errUserNotFound
	return
}
