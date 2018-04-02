FROM golang:1.8.1-alpine
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN go get github.com/hashicorp/consul

ENV SOURCES /go/src/github.com/johnantonusmaximus/cloud-go
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build 

WORKDIR ${SOURCES}
CMD ${SOURCES}/discovery/Consul/server/server
EXPOSE 8080