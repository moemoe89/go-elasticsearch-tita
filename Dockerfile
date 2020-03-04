FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/practicing-elasticsearch-golang

WORKDIR /go/src/github.com/moemoe89/practicing-elasticsearch-golang

COPY . /go/src/github.com/moemoe89/practicing-elasticsearch-golang

RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go mod download
RUN go install

ENTRYPOINT /go/bin/practicing-elasticsearch-golang

