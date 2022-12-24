package main

import (
	"web-api-gin/controller"
	"web-api-gin/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group(docs.SwaggerInfo.BasePath)
	{
		eg := v1.Group("/albums")
		{
			eg.GET("", ctrl.GetAlbums)
			eg.GET("/:id", ctrl.GetAlbumByID)
			eg.POST("", ctrl.PostAlbums)
			eg.DELETE("/:id", ctrl.DeleteAlbumByID)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8080")
}
