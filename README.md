# grpc-k8s-lb
A Kubernetes gRPC hello wold wirg istio/envoy load balancing and a simple gRPC-web UI

# How to build it

Requires the Go toolchain, the [`protoc` compiler](https://github.com/google/protobuf) and the [`protoc-gen-go` plugin](https://github.com/golang/protobuf).

``` shell
go get github.com/EtienneDufresne/grpc-k8s-lb
make build
```

# How to run it

## Start Server
``` shell
grpc-k8s-lb -s
```

## Start Client
``` shell
grpc-k8s-lb
```

## References

[Intro to protobuf and gRPC](https://medium.com/@shijuvar/building-high-performance-apis-in-go-using-grpc-and-protocol-buffers-2eda5b80771b)

[IETF gRCP RFC](https://tools.ietf.org/html/draft-kumar-rtgwg-grpc-protocol-00)

[gRCP Load Balancing](https://kubernetes.io/blog/2018/11/07/grpc-load-balancing-on-kubernetes-without-tears/)

[Envoy proxy gRCP docs](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/grpc)

[HTTP2 vs Webscocket](https://www.infoq.com/articles/websocket-and-http2-coexist)

[IETF websocket RFC](https://tools.ietf.org/html/rfc6455)

[grcp-web](https://github.com/grpc/grpc-web)

[gRCP chat implemented in go](https://github.com/rodaine/grpc-chat)
