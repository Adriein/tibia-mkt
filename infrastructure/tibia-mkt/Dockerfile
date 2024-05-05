FROM golang:alpine AS build

WORKDIR /var/www/tibia-mkt

COPY . .
RUN go build -o /var/www/tibia-mkt/bin/tibia-mkt-binary cmd/tibia-mkt/main.go


FROM scratch

COPY --from=build /var/www/tibia-mkt/bin/tibia-mkt-binary /var/www/tibia-mkt/bin/tibia-mkt-binary

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4/migrate.linux-386.tar.gz | tar xvz

ENTRYPOINT ["/var/www/tibia-mkt/bin/tibia-mkt-binary"]