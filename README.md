[![Build Status](https://travis-ci.org/moemoe89/practicing-elasticsearch-golang.svg?branch=master)](https://travis-ci.org/moemoe89/practicing-elasticsearch-golang)
[![Coverage Status](https://coveralls.io/repos/github/moemoe89/practicing-elasticsearch-golang/badge.svg?branch=master)](https://coveralls.io/github/moemoe89/practicing-elasticsearch-golang?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/practicing-elasticsearch-golang)](https://goreportcard.com/report/github.com/moemoe89/practicing-elasticsearch-golang)

# PRACTICING-ELASTICSEARCH-GOLANG #

Practicing Elasticsearch Using Golang (Gin Gonic Framework) as Programming Language, Elasticsearch as Search Engine Based

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ practicing-elasticsearch-golang/
  |     |
  |     +--+ main.go
  |        + api/
  |        + routers/
  |        + ... any other source code
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Setup and Build

* Setup Golang <https://golang.org/>
* Setup Elasticsearch <https://www.elastic.co/>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> practicing-elasticsearch-golang
```

## Running Application
Make config file for local :
```
$ cp config-sample.json config.json
```
Change Elasticsearch address based on your config :
```
http://:9200
```
Build
```
$ go build
```
Run
```
$ go run main.go
```

## How to Run with Docker
Make config file for docker :
```
$ cp config-sample.json config.json
```
Change Elasticsearch address based on your docker config :
```
http://elasticsearch:9200
```
Build
```
$ docker-compose build
```
Run
```
$ docker-compose up
```
Stop
```
$ docker-compose down
```

## How to Run Unit Test
Run
```
$ go test ./...
```
Run with cover
```
$ go test ./... -cover
```
Run with HTML output
```
$ go test ./... -coverprofile=c.out && go tool cover -html=c.out
```


## License

MIT