package model

import (
	"regexp"
)

type User struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// Public Method
func (user User) Validation() error {
	switch {
	case len(user.UserName) == 0:
		return ErrUserNameNotInformed
	case len(user.Password) == 0:
		return ErrUserPasswordNotInformed
	case user.passwordIsValid() == false:
		return ErrUserPasswordInvalid
	default:
		return nil
	}
}

// Public Method
func (me User) Equals(user User) bool {
	isEquals := false
	if me.UserName == user.UserName && me.Password == user.Password {
		isEquals = true
	}
	return isEquals
}

// private method
func (user User) passwordIsValid() bool {
	pattern := regexp.MustCompile("^(.*?)(?:[A-Z])(.*[0-9])(.?[A-z]){0,}")
	return pattern.MatchString(user.Password)
}
