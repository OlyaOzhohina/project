FROM golang:1.17.1-alpine3.14

WORKDIR /parser

RUN apk add --update make

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
