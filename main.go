package main

import (
	"practicing-elasticsearch-golang/controllers"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

func main() {
	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))

	r.GET("/ping", controllers.Ping)
	r.GET("/elasticsearch/destination", controllers.ElasticsearchDestinationGet)
	r.GET("/elasticsearch/destination/:id", controllers.ElasticsearchDestinationDetail)
	r.POST("/elasticsearch/destination", controllers.ElasticsearchDestinationAdd)
	r.DELETE("/elasticsearch/destination/:id", controllers.ElasticsearchDestinationDelete)

	r.Run()
}
