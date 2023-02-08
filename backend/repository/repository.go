package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gerodp/simpleBlogApp/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	DB    *gorm.DB
	Users model.UserRepository
	Posts model.PostRepository
}

func buildConnectionString(user string, pass string, address string, databaseName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, address, databaseName)
}

func NewRepository(user string, pass string, address string, databaseName string) (*Repository, error) {
	dsn := buildConnectionString(user, pass, address, databaseName)

	logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	repo := &Repository{
		DB:    db,
		Users: NewUserRepository(db),
		Posts: NewPostRepository(db),
	}

	err = repo.Users.Migrate()

	if err != nil {
		panic("failed to migrate Users")
	}

	err = repo.Posts.Migrate()

	if err != nil {
		panic("failed to migrate Posts")
	}

	return repo, err

}
