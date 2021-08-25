FROM golang:1.13-alpine

RUN apk update && apk add vim git build-base

ADD . /go/src/app

WORKDIR /go/src/app

# go modではいらないので、go getで個別に
RUN go get -v github.com/rubenv/sql-migrate/...
RUN go get github.com/pilu/fresh