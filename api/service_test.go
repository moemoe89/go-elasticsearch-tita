//
//  Practicing Elasticsearch
//
//  Copyright © 2016. All rights reserved.
//

package api_test

import (
	ap "github.com/moemoe89/go-elasticsearch-tita/api"
	"github.com/moemoe89/go-elasticsearch-tita/api/api_struct/form"
	"github.com/moemoe89/go-elasticsearch-tita/api/mocks"
	"github.com/moemoe89/go-elasticsearch-tita/config"

	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceCreate(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	req := &form.DestinationForm{
		ID:   1,
		Name: "Momo",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Create", req).Return(req, nil).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.Create(req)

		assert.NoError(t, err)
		assert.NotNil(t, row)
		assert.Equal(t,0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Create", req).Return(nil, errors.New("Unexpected database error")).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.Create(req)

		assert.Error(t, err)
		assert.Nil(t, row)
		assert.Equal(t,http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceFind(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Find", "").Return(nil, nil).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.Find("")

		assert.NoError(t, err)
		assert.Nil(t, row)
		assert.Equal(t,0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Find", "").Return(nil, errors.New("Unexpected database error")).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.Find("")

		assert.Error(t, err)
		assert.Nil(t, row)
		assert.Equal(t,http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}

func TestServiceFindByID(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("FindByID", "1").Return(nil, nil).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.FindByID("1")

		assert.NoError(t, err)
		assert.Nil(t, row)
		assert.Equal(t,0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("FindByID", "1").Return(nil, errors.New("Unexpected database error")).Once()
		u := ap.NewService(log, mockRepo)

		row, status, err := u.FindByID("1")

		assert.Error(t, err)
		assert.Nil(t, row)
		assert.Equal(t,http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}


func TestServiceFindIndex(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("FindIndex").Return(true, nil).Once()
		u := ap.NewService(log, mockRepo)

		exists, status, err := u.FindIndex()

		assert.NoError(t, err)
		assert.Equal(t, true, exists)
		assert.Equal(t,0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("FindIndex").Return(false, errors.New("Unexpected database error")).Once()
		u := ap.NewService(log, mockRepo)

		exists, status, err := u.FindIndex()

		assert.Error(t, err)
		assert.Equal(t, false, exists)
		assert.Equal(t,http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}


func TestServiceDelete(t *testing.T) {
	log := config.InitLog()
	mockRepo := new(mocks.Repository)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Delete", "1").Return(nil).Once()
		u := ap.NewService(log, mockRepo)

		status, err := u.Delete("1")

		assert.NoError(t, err)
		assert.Equal(t,0, status)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Delete", "1").Return(errors.New("Unexpected database error")).Once()
		u := ap.NewService(log, mockRepo)

		status, err := u.Delete("1")

		assert.Error(t, err)
		assert.Equal(t,http.StatusInternalServerError, status)

		mockRepo.AssertExpectations(t)
	})
}
