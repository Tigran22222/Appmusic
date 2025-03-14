package service

import (
	"appmusic"
	"appmusic/pkg/repository"
)

type AppmusicAlbumService struct {
	repo repository.Album
}

func NewAppmusicAlbumService(repo repository.Album) *AppmusicAlbumService {
	return &AppmusicAlbumService{repo: repo}
}

func (s *AppmusicAlbumService) Create(userId int, album appmusic.Album) (int, error) {
	return s.repo.Create(userId, album)
}

func (s *AppmusicAlbumService) GetAll(userId int) ([]appmusic.Album, error) {
	return s.repo.GetAll(userId)
}

func (s *AppmusicAlbumService) GetById(userId int, albumId int) (appmusic.Album, error) {
	return s.repo.GetById(userId, albumId)
}

func (s *AppmusicAlbumService) Delete(userId, albumId int) error {
	return s.repo.Delete(userId, albumId)
}

func (s *AppmusicAlbumService) Update(userId int, albumId int, input appmusic.UpdateAlbumInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, albumId, input)
}
