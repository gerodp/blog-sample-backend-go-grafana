package service

import (
	"errors"

	"github.com/gerodp/simpleBlogApp/model"
)

type PostService struct {
	postRepository model.PostRepository
}

func NewPostService(postRepo model.PostRepository) *PostService {
	return &PostService{
		postRepository: postRepo,
	}
}

func (u *PostService) Validate(post *model.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The title of the post is empty")
		return err
	}
	if post.Body == "" {
		err := errors.New("The body of the post is empty")
		return err
	}

	if post.AuthorID <= 0 {
		err := errors.New("The Author of the post is not valid.")
		return err
	}
	return nil
}

func (u *PostService) Create(post *model.Post) (*model.Post, error) {
	err := u.Validate(post)

	if err != nil {
		return nil, err
	} else {
		return u.postRepository.Save(post)
	}
}

func (u *PostService) Delete(post *model.Post) error {
	return u.postRepository.Delete(post)
}

func (u *PostService) Find(conds ...interface{}) ([]model.Post, error) {
	return u.postRepository.Find(conds...)
}
