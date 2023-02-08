package model

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type user User // prevent recursion
	x := user(u)
	x.Password = ""
	return json.Marshal(x)
}

type UserRepository interface {
	Save(user *User) (*User, error)

	FindByUsername(username string) (*User, error)

	Find(conds ...interface{}) ([]User, error)

	Delete(user *User) error

	Migrate() error
}

type UserService interface {
	Validate(user *User) error

	Create(user *User) (*User, error)

	FindUser(username string, password string) (*User, error)

	Delete(user *User) error

	Find(conds ...interface{}) ([]User, error)
}

func NewUser(username string, email string, password string) (u *User) {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
	}
}
