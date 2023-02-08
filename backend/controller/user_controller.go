package controller

import (
	"log"
	"net/http"

	"github.com/gerodp/simpleBlogApp/model"
	"github.com/gerodp/simpleBlogApp/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService model.UserService
}

func NewUserController(userRepository model.UserRepository) *UserController {
	return &UserController{
		userService: service.NewUserService(userRepository),
	}
}

func (u *UserController) Find(c *gin.Context) {
	users, err := u.userService.Find()

	if err != nil {
		log.Fatalf("Error retrieving users %s\n", err)
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Fatalf("Error getting parameters to create user %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err1 := u.userService.Create(&user)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
