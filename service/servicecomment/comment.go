package servicecomment

import (
	"log"

	"finall/entity"
	"finall/model/modelcomment"
	"finall/repository/repositorycomment"
	"finall/validation"

	"github.com/jinzhu/copier"
)

type ServiceComment interface {
	Create(request modelcomment.Request) (modelcomment.ResponseInsert, error)
	Update(request modelcomment.RequestUpdate, commentID uint) (modelcomment.ResponseUpdate, error)
	Delete(commentID uint) error
	Get() ([]modelcomment.Response, error)
}

type service struct {
	repo repositorycomment.RepositoryComment
}

func (s *service) Get() ([]modelcomment.Response, error) {
	comments, err := s.repo.Get()
	if err != nil {
		return []modelcomment.Response{}, err
	}

	var response []modelcomment.Response

	for _, comment := range comments {
		var singleResponse modelcomment.Response
		copier.Copy(&singleResponse, &comment)
		response = append(response, singleResponse)
	}

	return response, nil
}

func (s *service) Create(request modelcomment.Request) (modelcomment.ResponseInsert, error) {
	err := validation.ValidateComment(request)
	if err != nil {
		return modelcomment.ResponseInsert{}, err
	}

	var comment entity.Comment
	copier.Copy(&comment, &request)
	create, err := s.repo.Create(comment)
	if err != nil {
		return modelcomment.ResponseInsert{}, err
	}
	var response modelcomment.ResponseInsert
	copier.Copy(&response, &create)
	return response, nil
}

func (s *service) Update(request modelcomment.RequestUpdate, commentID uint) (modelcomment.ResponseUpdate, error) {
	err := validation.ValidateCommentUpdate(request)
	if err != nil {
		return modelcomment.ResponseUpdate{}, err
	}
	var comment entity.Comment
	copier.Copy(&comment, request)
	comment.ID = commentID
	update, err := s.repo.Update(comment)
	if err != nil {
		return modelcomment.ResponseUpdate{}, err
	}
	log.Println(update)
	var responseComment modelcomment.ResponseUpdate
	copier.Copy(&responseComment, update)
	return responseComment, nil
}

func (s *service) Delete(commentID uint) error {
	err := s.repo.Delete(commentID)
	if err != nil {
		return err
	}
	return nil
}

func New(repo repositorycomment.RepositoryComment) ServiceComment {
	return &service{repo: repo}
}
