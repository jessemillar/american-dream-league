FROM golang:alpine

RUN apk update
RUN apk upgrade
RUN apk add git

RUN mkdir -p /go/src/github.com/byuoitav
ADD . /go/src/github.com/byuoitav/av-api

WORKDIR /go/src/github.com/byuoitav/av-api
RUN go get -d -v
RUN go install -v

CMD ["/go/bin/american-dream-league"]

EXPOSE 8000
