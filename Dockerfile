FROM golang:alpine

MAINTAINER Pavle

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src/app

COPY . /go/src/app/

RUN go build .

EXPOSE $PORT

ENTRYPOINT ["./images"]