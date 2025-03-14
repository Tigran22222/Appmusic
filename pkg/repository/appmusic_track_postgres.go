package repository

import (
	"appmusic"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type AppmusicTrackPostgres struct {
	db *sqlx.DB
}

func NewAppmusicTrackPostgres(db *sqlx.DB) *AppmusicTrackPostgres {
	return &AppmusicTrackPostgres{db: db}
}

func (r *AppmusicTrackPostgres) Create(albumId int, track appmusic.Track) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var trackId int
	createTrackQuery := fmt.Sprintf("INSERT INTO tracks (title, duration, album_id) VALUES ($1, $2, $3) RETURNING id")

	row := tx.QueryRow(createTrackQuery, track.Title, track.Duration, albumId)
	err = row.Scan(&trackId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return trackId, nil
}

func (r *AppmusicTrackPostgres) GetAll(userId, albumId int) ([]appmusic.Track, error) {
	var tracks []appmusic.Track

	query := fmt.Sprintf(`
		SELECT t.* FROM %s t
		INNER JOIN %s a ON t.album_id = a.id
		WHERE t.album_id = $1 AND a.author_id = $2
	`, appmusicTracksTable, appmusicAlbumsTable)

	if err := r.db.Select(&tracks, query, albumId, userId); err != nil {
		return nil, err
	}

	return tracks, nil
}

func (r *AppmusicTrackPostgres) GetByIdFromAlbum(userId, albumId, trackId int) (appmusic.Track, error) {
	var track appmusic.Track
	query := `SELECT id, title, duration FROM tracks WHERE album_id = $1 AND id = $2`
	err := r.db.QueryRow(query, albumId, trackId).Scan(&track.Id, &track.Title, &track.Duration)
	if err != nil {
		return track, err
	}
	return track, nil
}

func (r *AppmusicTrackPostgres) DeleteByIdFromAlbum(userId, albumId, trackId int) error {

	query := fmt.Sprintf(`
		DELETE FROM %s t USING %s a
		WHERE t.id = $1 AND t.album_id = $2 AND a.id = t.album_id AND a.author_id = $3`,
		appmusicTracksTable, appmusicAlbumsTable)

	_, err := r.db.Exec(query, trackId, albumId, userId)
	return err
}

func (r *AppmusicTrackPostgres) UpdateByIdFromAlbum(userId, albumId, trackId int, input appmusic.UpdateTrackInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	// Если нет полей для обновления — выходим
	if len(setValues) == 0 {
		return nil
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		`UPDATE %s t SET %s FROM %s al 
		WHERE t.id = $%d AND al.id = $%d AND t.album_id = al.id AND al.author_id = $%d`,
		appmusicTracksTable, setQuery, appmusicAlbumsTable, argId, argId+1, argId+2,
	)

	args = append(args, trackId, albumId, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
