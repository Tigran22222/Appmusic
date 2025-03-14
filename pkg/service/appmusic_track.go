package service

import (
	"appmusic"
	"appmusic/pkg/repository"
)

type AppmusicTrackService struct {
	repo      repository.Track
	albumRepo repository.Album
}

func NewAppmusicTrackService(repo repository.Track, albumRepo repository.Album) *AppmusicTrackService {
	return &AppmusicTrackService{repo: repo, albumRepo: albumRepo}
}

func (s *AppmusicTrackService) Create(userId, albumId int, track appmusic.Track) (int, error) {
	// Проверяем, принадлежит ли альбом пользователю
	_, err := s.albumRepo.GetById(userId, albumId)
	if err != nil {
		return 0, err
	}

	// Создаем трек в указанном альбоме
	return s.repo.Create(albumId, track)
}

func (s *AppmusicTrackService) UpdateByIdFromAlbum(userId, albumId, trackId int, input appmusic.UpdateTrackInput) error {
	return s.repo.UpdateByIdFromAlbum(userId, albumId, trackId, input)
}

func (s *AppmusicTrackService) GetAll(userId, trackId int) ([]appmusic.Track, error) {
	return s.repo.GetAll(userId, trackId)
}

func (s *AppmusicTrackService) GetByIdFromAlbum(userId, albumId, trackId int) (appmusic.Track, error) {
	return s.repo.GetByIdFromAlbum(userId, albumId, trackId)
}

func (s *AppmusicTrackService) DeleteByIdFromAlbum(userId, albumId, trackId int) error {
	return s.repo.DeleteByIdFromAlbum(userId, albumId, trackId)
}
