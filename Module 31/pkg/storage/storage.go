package storage

import "time"

// Post - публикация.
type Post struct {
	ID         int    `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Content    string `json:"content" db:"content"`
	AuthorID   int    `json:"author_id" db:"author_id"`
	AuthorName string
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Posts() ([]Post, error) // получение всех публикаций
	AddPost(Post) error     // создание новой публикации
	UpdatePost(Post) error  // обновление публикации
	DeletePost(Post) error  // удаление публикации по ID
}
