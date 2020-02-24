//
//  Practicing Elasticsearch
//
//  Copyright © 2016. All rights reserved.
//

package config

import (
	"fmt"

	"gopkg.in/olivere/elastic.v5"
)

// InitElasticsearch will create a variable that represent the elastic.Client
func InitElasticsearch() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(Configuration.Elasticsearch.URL))
	if err != nil {
		return nil, fmt.Errorf("Failed to ping connection to elasticsearch: %s", err.Error())
	}

	return client, nil
}
