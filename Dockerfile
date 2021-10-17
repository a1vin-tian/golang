##
## Build
##

FROM golang:1.16-buster AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY httpserver ./httpserver

RUN ls; \
    go env -w GOPROXY="https://goproxy.io,direct"; \
    go mod vendor; \
    CGO_ENABLED=0 GOARCH=amd64 go build -o / ./httpserver


##
## Deploy
##
FROM scratch
WORKDIR /

COPY --from=build /httpserver /httpserver

EXPOSE 8080

ENV VERSION=1.0

LABEL lan="golang" app="httpserver"

ENTRYPOINT ["/httpserver"]


