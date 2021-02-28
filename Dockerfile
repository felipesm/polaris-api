# Dockerfile | Polaris API
FROM golang:1.16.0-alpine as build
LABEL maintainer='Felipe Maia | felipe.silwa@gmail.com'

WORKDIR $GOPATH/src/github.com/felipesm/polaris-api

COPY go.mod .

COPY *.go .

RUN go get github.com/felipesm/polaris-boleto

RUN go install github.com/felipesm/polaris-api

EXPOSE 3000

CMD [ "polaris-api" ]