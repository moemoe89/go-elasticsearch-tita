//
//  Practicing Elasticsearch
//
//  Copyright © 2016. All rights reserved.
//

package main

import (
	ap "github.com/moemoe89/practicing-elasticsearch-golang/api"
	conf "github.com/moemoe89/practicing-elasticsearch-golang/config"
	"github.com/moemoe89/practicing-elasticsearch-golang/routers"

	"fmt"

	"github.com/DeanThompson/ginpprof"
)

func main() {
	client, err := conf.InitElasticsearch()
	if err != nil {
		panic(err)
	}

	log := conf.InitLog()

	repo := ap.NewElasticsearchRepository(client)
	svc := ap.NewService(log, repo)

	app := routers.GetRouter(log, svc)
	ginpprof.Wrap(app)
	err = app.Run(":" + conf.Configuration.Port)
	if err != nil {
		panic(fmt.Sprintf("Can't start the app: %s", err.Error()))
	}
}
