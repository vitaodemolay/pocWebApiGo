package controller

import (
	"net/http"
	"web-api-gin/httputil"
	"web-api-gin/model"
	"web-api-gin/service"

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
