package service

import (
	"errors"
	"regexp"

	"github.com/gerodp/simpleBlogApp/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository model.UserRepository
}

func NewUserService(userRepo model.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepo,
	}
}

func validateEmail(email string) error {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)

	if re.MatchString(email) {
		return nil
	} else {
		return errors.New("The user has an invalid email address")
	}

}

func (u *UserService) Validate(user *model.User) error {
	if user == nil {
		err := errors.New("The user is empty")
		return err
	}
	if user.Username == "" {
		err := errors.New("The username of user is empty")
		return err
	}
	if user.Email == "" {
		err := errors.New("The email of user is empty")
		return err
	}
	emailErr := validateEmail(user.Email)

	if emailErr != nil {
		return emailErr
	}

	if user.Password == "" {
		err := errors.New("The password of user is empty")
		return err
	}
	return nil
}

func (u *UserService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *UserService) Create(user *model.User) (*model.User, error) {
	err := u.Validate(user)

	if err != nil {
		return nil, err
	} else {
		pass, err1 := u.hashPassword(user.Password)
		if err1 != nil {
			return nil, err1
		} else {
			user.Password = pass
			return u.userRepository.Save(user)
		}
	}
}

func (u *UserService) Delete(user *model.User) error {
	return u.userRepository.Delete(user)
}

func (u *UserService) FindUser(username string, password string) (*model.User, error) {
	user, err := u.userRepository.FindByUsername(username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	} else {
		if err1 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err1 != nil {
			return nil, err1
		} else {
			return user, nil
		}
	}
}

func (u *UserService) Find(conds ...interface{}) ([]model.User, error) {
	return u.userRepository.Find(conds...)
}
