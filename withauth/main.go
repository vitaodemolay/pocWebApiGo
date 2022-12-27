package main

import (
	"errors"
	"net/http"
	"web-api-gin/controller"
	"web-api-gin/docs"
	"web-api-gin/httputil"
	"web-api-gin/service"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var key = uuid.New()

const limitInSeconds float64 = 20 //20 seconds

// @title           Swagger Example API - Albums
// @version         1.0
// @description     This is a sample server albums.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support - VMRC
// @contact.url    http://www.swagger.io/support
// @contact.email  vitor.marcos@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

// @BasePath /api/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				    Description for what is this security definition being used
func main() {
	router := gin.Default()
	ctrl := controller.NewController()

	secret := []byte(key.String())

	router.Use(sessions.Sessions("Authentication", sessions.NewCookieStore(secret)))

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group(docs.SwaggerInfo.BasePath)
	{
		albums := v1.Group("/albums")
		{
			albums.GET("", ctrl.GetAlbums)
			albums.GET("/:id", ctrl.GetAlbumByID)
			albums.POST("", ctrl.PostAlbums)
			albums.DELETE("/:id", ctrl.DeleteAlbumByID)
		}

		users := v1.Group("/users")
		{
			users.Use(AuthRequired)
			users.GET("", ctrl.GetUsers)
			users.POST("", ctrl.PostUsers)
		}

		v1.POST("/signin", ctrl.SignIn)
		v1.POST("/signout", ctrl.SignOut)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8080")
}

func AuthRequired(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) == 0 {
		httputil.NewError(ctx, http.StatusUnauthorized, errors.New("Authorization is required Header"))
		ctx.Abort()
	} else if err, login := service.GetLoginByToken(authHeader); err != nil {
		httputil.NewError(ctx, http.StatusUnauthorized, errors.New("Token not found"))
		ctx.Abort()
	} else if login.IsExpired(limitInSeconds) {
		service.RemoveLoginByToken(login.UserId)
		httputil.NewError(ctx, http.StatusUnauthorized, errors.New("Token is expired"))
		ctx.Abort()
	}
	ctx.Next()
}
