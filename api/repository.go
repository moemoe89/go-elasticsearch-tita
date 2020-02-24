//
//  Practicing Elasticsearch
//
//  Copyright © 2016. All rights reserved.
//

package api

import (
	"github.com/moemoe89/practicing-elasticsearch-golang/api/api_struct/form"

	"strconv"

	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

const (
	elasticSearchIndex = "destination"
	elasticSearchType = "destination"
)

// Repository represent the repositories
type Repository interface {
	Create(destination *form.DestinationForm) (interface{}, error)
	Find(search string) (interface{}, error)
	FindByID(id string) (interface{}, error)
	FindIndex() (bool, error)
	Delete(id string) error
}

type elasticsearchRepository struct {
	Client *elastic.Client
}

// NewElasticsearchRepository will create an object that represent the Repository interface
func NewElasticsearchRepository(Client *elastic.Client) Repository {
	return &elasticsearchRepository{Client}
}

func (e *elasticsearchRepository) Create(destination *form.DestinationForm) (interface{}, error) {
	exists, err := e.Client.IndexExists(elasticSearchIndex).Do(context.TODO())
	if !exists {
		mapping := `
		{
			"settings" : {
				"number_of_shards" : 1
			},
			"mappings" : {
				"destination" : {
					"properties": {
						"name": {
								"type": "text",
								"fielddata": true
							}
						}
					}
				}
			}
		}`

		_, err := e.Client.CreateIndex(elasticSearchIndex).BodyString(mapping).Do(context.TODO())
		if err != nil {
			return nil, err
		}
	}

	put, err := e.Client.Index().Index(elasticSearchIndex).Type(elasticSearchType).Id(strconv.Itoa(destination.ID)).BodyJson(destination).Do(context.TODO())

	return put, err
}

func (e *elasticsearchRepository) Find(search string) (interface{}, error) {
	termQuery := elastic.NewTermQuery("name", search)
	searchResult, err := e.Client.Search().Index(elasticSearchIndex).Query(termQuery).Sort("name", true).From(0).Size(10).Pretty(true).Do(context.TODO())
	return searchResult, err
}

func (e *elasticsearchRepository) FindByID(id string) (interface{}, error) {
	get, err := e.Client.Get().Index(elasticSearchIndex).Type(elasticSearchType).Id(id).Do(context.TODO())
	return get, err
}

func (e *elasticsearchRepository) FindIndex() (bool, error) {
	exists, err := e.Client.IndexExists(elasticSearchIndex).Do(context.TODO())
	return exists, err
}

func (e *elasticsearchRepository) Delete(id string) error {
	_, err := e.Client.Delete().Index(elasticSearchIndex).Type(elasticSearchType).Id(id).Do(context.TODO())
	return err
}
