package service

import (
	"errors"
	"web-api-gin/model"
)

var errAlbumNotFound = errors.New("album not found")

// albums slice to seed record album data.
var dbAlbums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAllAlbums() []model.Album {
	return dbAlbums
}

func AddNewAlbum(newAlbum model.Album) {
	dbAlbums = append(dbAlbums, newAlbum)
}

func GetAlbumByID(id string) (err error, album model.Album) {
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range dbAlbums {
		if a.ID == id {
			album = a
			err = nil
			return
		}
	}
	err = errAlbumNotFound
	return
}

func RemoveAlbumById(id string) error {
	for idx, a := range dbAlbums {
		if a.ID == id {
			dbAlbums = append(dbAlbums[:idx], dbAlbums[idx+1:]...)
			return nil
		}
	}

	return errAlbumNotFound
}
