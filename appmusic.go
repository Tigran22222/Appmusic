package appmusic

// Album - структура для музыкального альбома
type Album struct {
	Id       int    `json:"id" db:"id"`
	AuthorId int    `json:"author_id" db:"author_id"`
	Title    string `json:"title" db:"title" binding:"required"`
}

// Track - структура для музыкального трека
type Track struct {
	Id       int    `json:"id" db:"id"`
	AlbumId  int    `json:"album_id" db:"album_id"`
	Title    string `json:"title" db:"title" binding:"required"`
	Duration int    `json:"duration" db:"duration"` // Длительность в секундах
}
