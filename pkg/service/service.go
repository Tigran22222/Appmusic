package service

import (
	"appmusic"
	"appmusic/pkg/repository"
)

type Authorization interface {
	GenerateToken(username, password string) (string, error)
	CreateUser(user appmusic.User) (int, error)
	ParseToken(token string) (int, error)
}

type Album interface {
	Create(userId int, album appmusic.Album) (int, error)
	GetAll(userId int) ([]appmusic.Album, error)
	GetById(userId int, albumId int) (appmusic.Album, error)
	Delete(userId, albumId int) error
	Update(userId, albumId int, input appmusic.UpdateAlbumInput) error
}

type Track interface {
	Create(userId, albumId int, track appmusic.Track) (int, error)
	GetAll(userId, trackId int) ([]appmusic.Track, error)
	GetByIdFromAlbum(userId, albumId, trackId int) (appmusic.Track, error)
	DeleteByIdFromAlbum(userId, albumId, trackId int) error
	UpdateByIdFromAlbum(userId, albumId, trackId int, input appmusic.UpdateTrackInput) error
}

type Service struct {
	Authorization
	Album
	Track
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Album:         NewAppmusicAlbumService(repos.Album),
		Track:         NewAppmusicTrackService(repos.Track, repos.Album),
	}
}
