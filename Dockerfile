FROM golang:1.14.3-stretch

RUN go get github.com/cixtor/rssfeed

ENTRYPOINT ["/go/bin/rssfeed"]

EXPOSE 9628
