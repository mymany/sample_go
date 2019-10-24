FROM golang:1.13.3-stretch

RUN go get "github.com/gin-gonic/gin"
RUN go get -u "github.com/golang/dep/cmd/dep"
RUN go get -u "github.com/derekparker/delve/cmd/dlv"
RUN go get -u "github.com/stretchr/testify"
RUN go get -u "github.com/stretchr/testify/assert"

WORKDIR /go/src/app
ADD . /go/src/app