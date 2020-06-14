//
//  Practicing Elasticsearch
//
//  Copyright © 2016. All rights reserved.
//

package api

import (
	"github.com/moemoe89/go-elasticsearch-tita/api/api_struct/form"

	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Service represent the services
type Service interface {
	Create(user *form.DestinationForm) (interface{}, int, error)
	Find(search string) (interface{}, int, error)
	FindByID(id string) (interface{}, int, error)
	FindIndex() (bool, int, error)
	Delete(id string) (int, error)
}

type implService struct {
	log        *logrus.Entry
	repository Repository
}

// NewService will create an object that represent the Service interface
func NewService(log *logrus.Entry, r Repository) Service {
	return &implService{log: log, repository: r}
}

func (s *implService) Create(req *form.DestinationForm) (interface{}, int, error) {
	user, err := s.repository.Create(req)
	if err != nil {
		s.log.Errorf("can't create data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return user, 0, nil
}

func (s *implService) Find(search string) (interface{}, int, error) {
	users, err := s.repository.Find(search)
	if err != nil {
		s.log.Errorf("can't get data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return users, 0, nil
}

func (s *implService) FindByID(id string) (interface{}, int, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		s.log.Errorf("can't get detail data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return user, 0, nil
}

func (s *implService) FindIndex() (bool, int, error) {
	exists, err := s.repository.FindIndex()
	if err != nil {
		s.log.Errorf("can't get index data: %s", err.Error())
		return exists, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return exists, 0, nil
}

func (s *implService) Delete(key string) (int, error) {
	err := s.repository.Delete(key)
	if err != nil {
		s.log.Errorf("can't delete data: %s", err.Error())
		return http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return 0, nil
}
