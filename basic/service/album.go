package service

import (
	"errors"
	"web-api-gin/model"
)

var ErrAlbumNotFound = errors.New("album not found")

// albums slice to seed record album data.
var database = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAllAlbums() []model.Album {
	return database
}

func AddNewAlbum(newAlbum model.Album) {
	database = append(database, newAlbum)
}

func GetAlbumByID(id string) (error, model.Album) {
	var foundAlbum model.Album

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range database {
		if a.ID == id {
			return nil, a
		}
	}
	return ErrAlbumNotFound, foundAlbum
}

func RemoveAlbumById(id string) error {
	for idx, a := range database {
		if a.ID == id {
			database = append(database[:idx], database[idx+1:]...)
			return nil
		}
	}

	return ErrAlbumNotFound
}
