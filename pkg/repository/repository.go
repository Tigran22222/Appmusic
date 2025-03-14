package repository

import (
	"appmusic"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user appmusic.User) (int, error)
	GetUser(username, password string) (appmusic.User, error)
}

type Album interface {
	Create(userId int, album appmusic.Album) (int, error)
	GetAll(userId int) ([]appmusic.Album, error)
	GetById(userId int, albumId int) (appmusic.Album, error)
	Delete(userId, albumId int) error
	Update(userId int, albumId int, input appmusic.UpdateAlbumInput) error
}

type Track interface {
	Create(albumId int, track appmusic.Track) (int, error)
	GetAll(userId, trackId int) ([]appmusic.Track, error)
	GetByIdFromAlbum(userId, albumId, trackId int) (appmusic.Track, error)
	DeleteByIdFromAlbum(userId, albumId, trackId int) error
	UpdateByIdFromAlbum(userId, albumId, trackId int, input appmusic.UpdateTrackInput) error
}

type Repository struct {
	Authorization
	Album
	Track
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Album:         NewAppmusicAlbumPostgres(db),
		Track:         NewAppmusicTrackPostgres(db),
	}
}
