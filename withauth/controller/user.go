package controller

import (
	"errors"
	"net/http"
	"strings"
	"web-api-gin/httputil"
	"web-api-gin/model"
	"web-api-gin/service"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary get users list
// @Description Get users responds with the list of all users as JSON.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} []model.User
// @Router /users [get]
func (controller *Controller) GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.GetAllUsers())
}

// PostUsers godoc
// @Summary post an user
// @Description adds an user from JSON received in the request body.
// @Tags users
// @Accept json
// @Produce json
// @Param        user body      model.User  true  "Add User"
// @Success 200 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Router /users [post]
func (controller *Controller) PostUsers(c *gin.Context) {
	var newUser model.User

	if err := c.BindJSON(&newUser); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	} else if err := newUser.Validation(); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	service.AddNewUser(newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// SignIn godoc
// @Summary post an user login
// @Description logon the user.
// @Tags users
// @Accept x-www-form-urlencoded
// @Produce json
// @Param        authetication	 formData	model.User	false  "username and password in formData"
// @Success 200
// @Failure 400 {object} httputil.HTTPError
// @Failure 401 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/signin [post]
func (controller *Controller) SignIn(c *gin.Context) {
	session := sessions.Default(c)
	userName := c.PostForm("userName")
	password := c.PostForm("password")

	if strings.Trim(userName, " ") == "" || strings.Trim(password, " ") == "" {
		httputil.NewError(c, http.StatusBadRequest, errors.New("Parameters can't be empty"))
		return
	}

	var login model.User
	login.UserName = userName
	login.Password = password

	if err, user := service.GetUserByUserName(userName); err != nil {
		httputil.NewError(c, http.StatusUnauthorized, err)
		return
	} else if !login.Equals(user) {
		httputil.NewError(c, http.StatusUnauthorized, errors.New("user name or password is invalid"))
		return
	}

	session.Set(userName, userName)
	if err := session.Save(); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, errors.New("Failed to save session"))
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}
