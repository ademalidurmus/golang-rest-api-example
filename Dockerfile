FROM golang:1.15

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
ADD ./app /go/src/app

RUN go get -v