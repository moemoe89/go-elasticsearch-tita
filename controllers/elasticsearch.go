package controllers

import (
	"practicing-elasticsearch-golang/models"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	elastic "gopkg.in/olivere/elastic.v5"
	"golang.org/x/net/context"
)

const(
	elasticSearchIndex = "destination"
	elasticSearchType  = "destination"
	elasticSearchURL   = "http://127.0.0.1:9200"
)

var client *elastic.Client
var err error

func init(){
	client, err = elastic.NewClient(elastic.SetURL(elasticSearchURL))
	if err != nil {
		panic(err)
	}
}

func ElasticsearchDestinationDetail(c *gin.Context){

	id := c.Param("id")
	exists, err := client.IndexExists(elasticSearchIndex).Do(context.TODO())
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}
	if !exists {
		JSONResponse(c, http.StatusNotFound, "Index not exists.", false)
		return
	}

	// Get tweet with specified ID
	get, err := client.Get().
	Index(elasticSearchIndex).
	Type(elasticSearchType).
	Id(id).
	Do(context.TODO())
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}
	if get.Found {
		JSONResponseData(c, http.StatusOK, "Get destination.", true,get)
		return
	} else {
		JSONResponse(c, http.StatusNotFound, "Destination not found", false)
		return
	}

}

func ElasticsearchDestinationGet(c *gin.Context){

	search := c.Query("search")

	exists, err := client.IndexExists(elasticSearchIndex).Do(context.TODO())
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}
	if !exists {
		JSONResponse(c, http.StatusNotFound, "Index not exists.", false)
		return
	}

	termQuery := elastic.NewTermQuery("name", search)
	searchResult, err := client.Search().
	Index(elasticSearchIndex).
	Query(termQuery).
	Sort("name", true).
	From(0).Size(10).
	Pretty(true).
	Do(context.TODO())
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}

	JSONResponseData(c, http.StatusOK, "Get destinations.", true,searchResult)
}

func ElasticsearchDestinationAdd(c *gin.Context){

	exists, err := client.IndexExists(elasticSearchIndex).Do(context.TODO())
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}
	if !exists {
		// Create a new index.
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

		createIndex, err := client.CreateIndex(elasticSearchIndex).BodyString(mapping).Do(context.TODO())
		if err != nil {
			JSONResponse(c, http.StatusInternalServerError, err.Error(), false)
			return
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	destination := models.Destination{}
	err = c.BindJSON(&destination)
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}

	put, err := client.Index().
	Index(elasticSearchIndex).
	Type(elasticSearchType).
	Id(strconv.Itoa(destination.ID)).
	BodyJson(destination).
	Do(context.TODO())
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}

	JSONResponseData(c, http.StatusOK, "Success add destination.", true,put)
}

func ElasticsearchDestinationDelete(c *gin.Context){

	id := c.Param("id")

	exists, err := client.IndexExists(elasticSearchIndex).Do(context.TODO())
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}
	if !exists {
		JSONResponse(c, http.StatusNotFound, "Index not exists.", false)
		return
	}

	res, err := client.Delete().
	Index(elasticSearchIndex).
	Type(elasticSearchType).
	Id(id).
	Do(context.TODO())
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}
	if res.Found {
		JSONResponse(c, http.StatusOK, "Success delete data.", true)
		return
	} else {
		JSONResponse(c, http.StatusNotFound, "Destination not found", false)
		return
	}

}

func JSONResponse(c *gin.Context, httpStatus int, message string, status bool){
	c.IndentedJSON(httpStatus, gin.H{
		"message": message,
		"status":  status,
	})
}

func JSONResponseData(c *gin.Context, httpStatus int, message string, status bool, data interface{}){
	c.IndentedJSON(httpStatus, gin.H{
		"data": data,
		"message": message,
		"status": status,
	})
}
