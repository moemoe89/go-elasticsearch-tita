//
//  Practicing Elasticsearch
//
//  Copyright © 2016. All rights reserved.
//

package routers

import (
	ap "github.com/moemoe89/practicing-elasticsearch-golang/api"
	mw "github.com/moemoe89/practicing-elasticsearch-golang/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetRouter will create a variable that represent the gin.Engine
func GetRouter(log *logrus.Entry, svc ap.Service) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(mw.CORS)
	r.GET("/", ap.Ping)
	r.GET("/ping", ap.Ping)

	ctrl := ap.NewCtrl(log, svc)

	r.GET("/elasticsearch/destination", ctrl.Find)
	r.GET("/elasticsearch/destination/:id", ctrl.FindByID)
	r.POST("/elasticsearch/destination", ctrl.Create)
	r.DELETE("/elasticsearch/destination/:id", ctrl.Delete)

	return r
}
