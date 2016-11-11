FROM golang:alpine

RUN apk update
RUN apk upgrade
RUN apk add git

RUN mkdir -p /go/src/github.com/jessemillar
ADD . /go/src/github.com/jessemillar/american-dream-league

WORKDIR /go/src/github.com/jessemillar/american-dream-league
RUN go get -d -v
RUN go install -v

CMD ["/go/bin/american-dream-league"]

EXPOSE 9999
