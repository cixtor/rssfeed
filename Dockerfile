FROM golang:1.10.4-stretch

ADD . /go/src/github.com/cixtor/rssfeed

RUN go get github.com/cixtor/middleware

RUN go install github.com/cixtor/rssfeed

ENTRYPOINT ["/go/bin/rssfeed"]

EXPOSE 9628
