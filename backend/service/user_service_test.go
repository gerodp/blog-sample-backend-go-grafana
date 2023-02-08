package service

import (
	"testing"

	"github.com/gerodp/simpleBlogApp/model"
)

type UserRepositoryStub struct{}

func (u UserRepositoryStub) Save(user *model.User) (*model.User, error) {
	return &model.User{}, nil
}

func (u UserRepositoryStub) Find() ([]model.User, error) {
	var users []model.User
	return users, nil
}

func (u UserRepositoryStub) Delete(user *model.User) error {
	return nil
}

func (u UserRepositoryStub) Migrate() error {
	return nil
}

func Test_Validate(t *testing.T) {

	uServ := NewUserService(UserRepositoryStub{})

	err := uServ.Validate(nil)

	if err == nil || err.Error() != "The user is empty" {
		t.Errorf("Expected to fail with '%s' but got '%s'", "The user is empty", err)
	}

	err1 := uServ.Validate(&model.User{})

	if err1 == nil || err1.Error() != "The username of user is empty" {
		t.Errorf("Expected to fail with '%s' but got '%s'", "The username of user is empty", err1)
	}

	err2 := uServ.Validate(&model.User{Username: "p"})

	if err2 == nil || err2.Error() != "The email of user is empty" {
		t.Errorf("Expected to fail with '%s' but got '%s'", "The email of user is empty", err2)
	}

	err22 := uServ.Validate(&model.User{Username: "p", Email: "bbb@"})

	if err22 == nil || err22.Error() != "The user has an invalid email address" {
		t.Errorf("Expected to fail with '%s' but got '%s'", "The user has an invalid email address", err22)
	}

	err3 := uServ.Validate(&model.User{Username: "p", Email: "gg@bb.es"})

	if err3 == nil || err3.Error() != "The password of user is empty" {
		t.Errorf("Expected to fail with '%s' but got '%s'", "The password of user is empty", err3)
	}

	err4 := uServ.Validate(&model.User{Username: "p", Email: "gg@bb.es", Password: "112"})

	if err4 != nil {
		t.Errorf("Expected to succeed but got '%s'", err4)
	}

}
