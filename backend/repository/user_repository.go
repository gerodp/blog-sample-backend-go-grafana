package repository

import (
	"github.com/gerodp/simpleBlogApp/model"
	"gorm.io/gorm"
)

// Implements the models.UserRepository interface
type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Save(user *model.User) (*model.User, error) {
	return user, u.DB.Create(user).Error
}

func (u *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := u.DB.First(&user, "username = ?", username).Error

	return &user, err
}

func (u *UserRepository) Find(conds ...interface{}) ([]model.User, error) {
	var users []model.User
	err := u.DB.Find(&users, conds...).Error

	return users, err
}

func (u *UserRepository) Delete(user *model.User) error {
	return u.DB.Delete(user).Error
}

func (u *UserRepository) Migrate() error {
	return u.DB.AutoMigrate(&model.User{})
}
