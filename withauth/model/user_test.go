package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserEquals(t *testing.T) {
	//Arrage
	userA := User{
		UserName: "admin",
		Password: "Passw0rd",
	}
	userB := userA

	//Act
	result := userA.Equals(userB)

	//Assert
	assert.True(t, result)
}

func TestUserPasswordIsValid(t *testing.T) {
	//Arrange
	userA := User{
		UserName: "admin",
		Password: "Passw0rd",
	}
	//Act
	result := userA.passwordIsValid()

	//Assert
	assert.True(t, result)
}
