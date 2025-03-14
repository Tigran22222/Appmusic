package appmusic

import (
	"errors"
	"time"
)

// Album - структура для музыкального альбома
type Album struct {
	Id        int       `json:"id" db:"id"`
	AuthorId  int       `json:"author_id" db:"author_id"`
	Title     string    `json:"title" db:"title" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Track - структура для музыкального трека
type Track struct {
	Id       int    `json:"id" db:"id"`
	AlbumId  int    `json:"album_id" db:"album_id"`
	Title    string `json:"title" db:"title" binding:"required"`
	Duration int    `json:"duration" db:"duration"` // Длительность в секундах
}

type UpdateAlbumInput struct {
	Title *string `json:"title" binding:"required"`
	Done  *bool   `json:"done" binding:"required"`
}

func (i UpdateAlbumInput) Validate() error {
	if i.Title == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type UpdateTrackInput struct {
	Title *string `json:"title" binding:"required"`
	Done  *bool   `json:"done" binding:"required"`
}

func (i UpdateTrackInput) Validate() error {
	if i.Title == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
