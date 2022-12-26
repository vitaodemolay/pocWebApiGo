package main

import (
	"web-api-gin/controller"
	"web-api-gin/docs"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var key = uuid.New()

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
			users.GET("", ctrl.GetUsers)
			users.POST("", ctrl.PostUsers)
			users.POST("/signin", ctrl.SignIn)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8080")
}
