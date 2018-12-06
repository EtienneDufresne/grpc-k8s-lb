FROM golang:1.11-alpine as builder
RUN apk add --no-cache bash make git
RUN go get -d \
    github.com/pkg/errors \
    golang.org/x/net/context \
    google.golang.org/grpc \
    github.com/golang/protobuf/ptypes
WORKDIR /go/src/github.com/EtienneDufresne/grpc-k8s-lb
COPY . /go/src/github.com/EtienneDufresne/grpc-k8s-lb
RUN make build

FROM alpine:latest
EXPOSE 8080
COPY --from=builder /go/src/github.com/EtienneDufresne/grpc-k8s-lb/grpc-k8s-lb /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/grpc-k8s-lb && chmod +x /usr/local/bin/grpc-k8s-lb
USER nobody
ENTRYPOINT ["grpc-k8s-lb", "-s"]
