FROM golang:alpine AS build

WORKDIR /var/www/tibia-mkt

COPY ../../server .

RUN go mod download

RUN go build -o bin/tibia-mkt-bin cmd/tibia-mkt/main.go

RUN apk update && apk add curl tar bash

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz


FROM golang:alpine

WORKDIR /var/www/tibia-mkt

COPY --from=build /var/www/tibia-mkt/bin/tibia-mkt-bin /var/www/tibia-mkt/bin/tibia-mkt-bin

ENTRYPOINT ["/var/www/tibia-mkt/bin/tibia-mkt-bin"]