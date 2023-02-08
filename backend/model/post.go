package model

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AuthorID  uint      `json:"author_id"`
	Author    User      `gorm:"constraint:OnDelete:CASCADE;"`
	Title     string    `json:"title"`
	Body      string    `json:"body" gorm:"type:text"`
}

type PostRepository interface {
	Save(post *Post) (*Post, error)

	Find(conds ...interface{}) ([]Post, error)

	Delete(post *Post) error

	Migrate() error
}

type PostService interface {
	Validate(post *Post) error

	Create(post *Post) (*Post, error)

	Delete(post *Post) error

	Find(conds ...interface{}) ([]Post, error)
}
