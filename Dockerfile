FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/go-elasticsearch-tita

WORKDIR /go/src/github.com/moemoe89/go-elasticsearch-tita

COPY . /go/src/github.com/moemoe89/go-elasticsearch-tita

RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go mod download
RUN go install

ENTRYPOINT /go/bin/go-elasticsearch-tita
