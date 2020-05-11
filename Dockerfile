# Compile stage
FROM golang:1.14.2 AS build-env

EXPOSE 40000

# Build Delve
RUN go get github.com/go-delve/delve/cmd/dlv

ADD . /dockerdev
WORKDIR /dockerdev

CMD ["/go/bin/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "test", "./..."]
