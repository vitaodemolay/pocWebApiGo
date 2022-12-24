package model

import "errors"

var (
	ErrNoRow                  = errors.New("no rows in result set")
	ErrAlbumIdNotInformed     = errors.New("ID field is required")
	ErrAlbumTitleNotInformed  = errors.New("Title field is required")
	ErrAlbumArtistNotInformed = errors.New("Artist field is required")
	ErrAlbumPriceIsInvalid    = errors.New("Price field must be greater than zero")
)
