//
//  Practicing Elasticsearch
//
//  Copyright © 2016. All rights reserved.
//

package api_test

import (
	"github.com/moemoe89/go-elasticsearch-tita/api/api_struct/form"
	"github.com/moemoe89/go-elasticsearch-tita/api/mocks"
	"github.com/moemoe89/go-elasticsearch-tita/config"
	"github.com/moemoe89/go-elasticsearch-tita/routers"

	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessCreate(t *testing.T) {
	destForm := &form.DestinationForm{
		ID:   1,
		Name: "name",
	}

	j, err := json.Marshal(destForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("FindIndex").Return(true, 0, nil)
	mockService.On("Create", destForm).Return(nil, http.StatusCreated, nil)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/elasticsearch/destination", strings.NewReader(string(j)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFailCreate(t *testing.T) {
	destForm := &form.DestinationForm{
		ID:   1,
		Name: "name",
	}

	j, err := json.Marshal(destForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("FindIndex").Return(true, 0, nil)
	mockService.On("Create", destForm).Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/elasticsearch/destination", strings.NewReader(string(j)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFailBindCreate(t *testing.T) {
	destForm := map[string]interface{} {
		"id": "a",
	}

	j, err := json.Marshal(destForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/elasticsearch/destination", strings.NewReader(string(j)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFailCreateFindIndex(t *testing.T) {
	destForm := &form.DestinationForm{
		ID:   1,
		Name: "name",
	}

	j, err := json.Marshal(destForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("FindIndex").Return(false, http.StatusInternalServerError, errors.New("Unexpected database error"))

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/elasticsearch/destination", strings.NewReader(string(j)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFalseCreateFindIndex(t *testing.T) {
	destForm := &form.DestinationForm{
		ID:   1,
		Name: "name",
	}

	j, err := json.Marshal(destForm)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	mockService.On("FindIndex").Return(false, http.StatusNotFound, nil)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/elasticsearch/destination", strings.NewReader(string(j)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestSuccessFind(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Find", "").Return(nil, 0, nil)
	mockService.On("FindIndex").Return(true, 0, nil)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/elasticsearch/destination", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFailFind(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Find", "").Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/elasticsearch/destination", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFailFindFindIndex(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Find", "").Return(nil, 0, nil)
	mockService.On("FindIndex").Return(false, http.StatusInternalServerError, errors.New("Unexpected database error"))

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/elasticsearch/destination", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFalseFindFindIndex(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Find", "").Return(nil, 0, nil)
	mockService.On("FindIndex").Return(false, http.StatusNotFound, nil)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/elasticsearch/destination", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestSuccessFindByID(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("FindByID", "1").Return(nil, 0, nil)
	mockService.On("FindIndex").Return(true, 0, nil)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/elasticsearch/destination/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFailFindByID(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("FindByID", "1").Return(nil, http.StatusInternalServerError, errors.New("Unexpected database error"))

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/elasticsearch/destination/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFailFindByIDFindIndex(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("FindByID", "1").Return(nil, 0, nil)
	mockService.On("FindIndex").Return(false, http.StatusInternalServerError, errors.New("Unexpected database error"))

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/elasticsearch/destination/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFalseFindByIDFindIndex(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("FindByID", "1").Return(nil, 0, nil)
	mockService.On("FindIndex").Return(false, http.StatusNotFound, nil)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/elasticsearch/destination/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}


func TestSuccessDelete(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Delete", "1").Return(0, nil)
	mockService.On("FindIndex").Return(true, 0, nil)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/elasticsearch/destination/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFailDelete(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Delete", "1").Return(http.StatusInternalServerError, errors.New("Unexpected database error"))

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/elasticsearch/destination/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFailDeleteFindIndex(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Delete", "1").Return(0, nil)
	mockService.On("FindIndex").Return(false, http.StatusInternalServerError, errors.New("Unexpected database error"))

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/elasticsearch/destination/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFalseDeleteFindIndex(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Delete", "1").Return(0, nil)
	mockService.On("FindIndex").Return(false, http.StatusNotFound, nil)

	log := config.InitLog()

	router := routers.GetRouter(log, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/elasticsearch/destination/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

