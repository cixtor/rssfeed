FROM golang:1.10.4-stretch

RUN go get github.com/cixtor/rssfeed

ENTRYPOINT ["/go/bin/rssfeed"]

EXPOSE 9628
