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

// SignIn godoc
// @Summary post an user login
// @Description logon the user.
// @Tags signin
// @Accept x-www-form-urlencoded
// @Produce json
// @Param        authetication	 formData	model.User	false  "username and password in formData"
// @Success 200
// @Failure 400 {object} httputil.HTTPError
// @Failure 401 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /signin [post]
func (controller *Controller) SignIn(c *gin.Context) {
	session := sessions.Default(c)
	userName := c.PostForm("userName")
	password := c.PostForm("password")

	if strings.Trim(userName, " ") == "" || strings.Trim(password, " ") == "" {
		httputil.NewError(c, http.StatusBadRequest, errors.New("Parameters can't be empty"))
		return
	}

	var userTologin model.User
	userTologin.UserName = userName
	userTologin.Password = password

	if err, user := service.GetUserByUserName(userName); err != nil {
		httputil.NewError(c, http.StatusUnauthorized, err)
		return
	} else if !userTologin.Equals(user) {
		httputil.NewError(c, http.StatusUnauthorized, errors.New("user name or password is invalid"))
		return
	}
	login := model.CreateLogin(userTologin)

	session.Set(userName, login.UserId)
	if err := session.Save(); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, errors.New("Failed to save session"))
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": login.UserId})
}

// SignOut godoc
// @Summary post an user logout
// @Description logout the user.
// @Tags signout
// @Accept x-www-form-urlencoded
// @Produce json
// @Param        authetication	 formData	model.User	false  "username and password in formData"
// @Success 200
// @Failure 400 {object} httputil.HTTPError
// @Failure 401 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /signout [post]
func (controller *Controller) SignOut(c *gin.Context) {
	session := sessions.Default(c)
	userName := c.PostForm("userName")

	if strings.Trim(userName, " ") == "" {
		httputil.NewError(c, http.StatusBadRequest, errors.New("Parameters can't be empty"))
		return
	}

	userId := session.Get(userName)
	if userId == nil {
		httputil.NewError(c, http.StatusBadRequest, errors.New("Invalid session token"))
		return
	}

	service.RemoveLoginByToken(userId.(string))
	session.Delete(userName)

	if err := session.Save(); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, errors.New("Failed to save session"))
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully sign out"})
}
