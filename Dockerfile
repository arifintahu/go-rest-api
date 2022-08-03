# syntax=docker/dockerfile:1

## Build
FROM golang:1.18-alpine as build

WORKDIR /
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build main.go

## Deploy
FROM alpine:3

RUN mkdir /app

WORKDIR /

COPY --from=build /main /app

COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh

EXPOSE 8080

ENTRYPOINT ["sh", "/app/start.sh"]
