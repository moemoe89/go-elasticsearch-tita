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
$ cp config-sample.json config-local.json
```
Build
```
$ go build
```
Run
```
$ go run main.go
```
