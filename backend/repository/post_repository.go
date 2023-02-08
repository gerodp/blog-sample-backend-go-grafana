package repository

import (
	"github.com/gerodp/simpleBlogApp/model"
	"github.com/gerodp/simpleBlogApp/repository/helper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Implements the models.PostRepository interface
type PostRepository struct {
	DB        *gorm.DB
	collector *helper.QueryMetricCollector
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		DB:        db,
		collector: helper.NewQueryCollector("post_repository"),
	}
}

func (u *PostRepository) Save(post *model.Post) (*model.Post, error) {
	st := u.collector.Start()
	p, err := post, u.DB.Create(post).Error
	u.collector.End("save", st)
	return p, err
}

func (u *PostRepository) Find(conds ...interface{}) ([]model.Post, error) {
	st := u.collector.Start()

	var posts []model.Post

	cleanedConds := helper.ClearPaginateParams(conds...)

	err := u.DB.Scopes(helper.Paginate(conds...)).Preload("Author").Find(&posts, cleanedConds...).Error

	u.collector.End("find", st)
	return posts, err
}

func (u *PostRepository) Delete(post *model.Post) error {
	st := u.collector.Start()

	err := u.DB.Delete(&model.Post{}, post.ID).Error

	u.collector.End("delete", st)

	return err
}

func (u *PostRepository) Migrate() error {
	return u.DB.Omit(clause.Associations).AutoMigrate(&model.Post{})
}
