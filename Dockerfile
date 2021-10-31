FROM golang:1.16.7-alpine3.14 as builder
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

WORKDIR 


RUN go mod download

RUN go build -o ConfigManager main.go

FROM alpine:3.14

