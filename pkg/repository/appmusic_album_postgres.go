package repository

import (
	"appmusic"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AppmusicAlbumPostgres struct {
	db *sqlx.DB
}

func NewAppmusicAlbumPostgres(db *sqlx.DB) *AppmusicAlbumPostgres {
	return &AppmusicAlbumPostgres{db: db}
}

func (r *AppmusicAlbumPostgres) Create(userId int, album appmusic.Album) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var albumId int
	createAlbumQuery := fmt.Sprintf("INSERT INTO %s (author_id, title) VALUES ($1, $2) RETURNING id", appmusicAlbumsTable)
	row := tx.QueryRow(createAlbumQuery, userId, album.Title)
	if err := row.Scan(&albumId); err != nil {
		tx.Rollback()
		return 0, err
	}

	//createUsersQuery := fmt.Sprintf("INSERT INTO %s (user_id, author_id) VALUES ($1, $2)", usersTable)
	createUsersQuery := fmt.Sprintf("INSERT INTO users_albums (user_id, album_id) VALUES ($1, $2)")
	_, err = tx.Exec(createUsersQuery, userId, albumId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return albumId, tx.Commit()
}

func (r *AppmusicAlbumPostgres) GetAll(userId int) ([]appmusic.Album, error) {
	var albums []appmusic.Album

	query := fmt.Sprintf(`
		SELECT albums.* 
		FROM %s albums 
		INNER JOIN users_albums ua ON albums.id = ua.album_id 
		WHERE ua.user_id = $1
	`, appmusicAlbumsTable) // users_albums напрямую в запросе

	err := r.db.Select(&albums, query, userId)
	if err != nil {
		return nil, err
	}

	return albums, nil
}

func (r *AppmusicAlbumPostgres) GetById(userId int, albumId int) (appmusic.Album, error) {
	var album appmusic.Album

	query := fmt.Sprintf(`
        SELECT albums.* 
        FROM %s albums 
        INNER JOIN users_albums ua ON albums.id = ua.album_id 
        WHERE ua.user_id = $1 AND ua.album_id = $2
    `, appmusicAlbumsTable)

	err := r.db.Get(&album, query, userId, albumId)

	return album, err
}

func (r *AppmusicAlbumPostgres) Delete(userId, albumId int) error {
	query := `
		DELETE FROM albums 
		WHERE id = (
			SELECT album_id FROM users_albums WHERE user_id = $1 AND album_id = $2
		)
	`
	fmt.Println("Generated SQL Query:", query)

	_, err := r.db.Exec(query, userId, albumId)
	return err
}

func (r *AppmusicAlbumPostgres) Update(userId int, albumId int, input appmusic.UpdateAlbumInput) error {

	query := fmt.Sprintf("UPDATE %s SET title=$1 WHERE id=$2 AND author_id=(SELECT author_id FROM %s WHERE id=$3)",
		appmusicAlbumsTable, usersTable,
	)
	if input.Title == nil {
		return errors.New("no title provided for update")
	}
	_, err := r.db.Exec(query, *input.Title, albumId, userId)
	return err
}
