package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gerodp/simpleBlogApp/controller/helper"
	"github.com/gerodp/simpleBlogApp/model"
	"github.com/gerodp/simpleBlogApp/service"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService model.PostService
}

func NewPostController(postRepository model.PostRepository) *PostController {
	return &PostController{
		postService: service.NewPostService(postRepository),
	}
}

func (u *PostController) Find(c *gin.Context) {

	aid, ok := c.GetQuery("author_id")

	var posts []model.Post
	var err error
	var conds []interface{}

	if ok {
		if _, err := strconv.ParseUint(aid, 10, 0); err == nil {
			conds = make([]interface{}, 2)
			conds[0] = "author_id = ?"
			conds[1] = aid
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	pParams, err := helper.ParsePaginationParams(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conds = append(conds, pParams...)

	posts, err = u.postService.Find(conds...)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (u *PostController) CreatePost(c *gin.Context) {
	var post model.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		log.Fatalf("Error getting parameters to create post %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPost, err1 := u.postService.Create(&post)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdPost)
}

func (p *PostController) DeletePost(c *gin.Context) {
	id := c.Param("id")
	var err error

	if id != "" {
		var cid uint64
		if cid, err = strconv.ParseUint(id, 10, 0); err == nil {
			post := model.Post{
				ID: uint(cid),
			}

			err = p.postService.Delete(&post)
			if err == nil {
				c.JSON(http.StatusFound, "")
				return
			}

		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing post id parameter"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
