package controller

import (
	"net/http"
	"web-api-gin/httputil"
	"web-api-gin/model"
	"web-api-gin/service"

	"github.com/gin-gonic/gin"
)

// GetAlbums godoc
// @Summary get albums list
// @Description Get Albums responds with the list of all albums as JSON.
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {object} []model.Album
// @Router /albums [get]
func (controller *Controller) GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.GetAllAlbums())
}

// PostAlbums godoc
// @Summary post an albums
// @Description adds an album from JSON received in the request body.
// @Tags albums
// @Accept json
// @Produce json
// @Param        album body      model.Album  true  "Add Album"
// model.Album
// @Success 200 {object} model.Album
// @Failure 400 {object} httputil.HTTPError
// @Router /albums [post]
func (controller *Controller) PostAlbums(c *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	} else if err := newAlbum.Validation(); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	// Add the new album to the slice.
	service.AddNewAlbum(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID godoc
// @Summary get an albums by id
// @Description Get the album as JSON whose ID value matches the id.
// @Tags albums
// @Accept json
// @Produce json
// @Param        id   path      int  true  "Album ID"
// @Success 200 {object} model.Album
// @Failure 404 {object} httputil.HTTPError
// @Router /albums/{id} [get]
func (controller *Controller) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	err, album := service.GetAlbumByID(id)
	if err == nil {
		c.IndentedJSON(http.StatusOK, album)
		return
	}

	httputil.NewError(c, http.StatusNotFound, err)
	//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// DeleteAlbumByID godoc
// @Summary delete an albums by id
// @Description Locates the album whose ID value matches the id and remove from collection.
// @Tags albums
// @Accept json
// @Produce json
// @Param        id   path      int  true  "Album ID"
// @Success 200
// @Failure 404 {object} httputil.HTTPError
// @Router /albums/{id} [delete]
func (controller *Controller) DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	if err := service.RemoveAlbumById(id); err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album has deleted"})
}
