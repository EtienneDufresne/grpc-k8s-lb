# grpc-k8s-lb
A Kubernetes gRPC echo client/server with istio/envoy load balancing.

# Requirements

- [go](https://golang.org/doc/install)
- [protoc compiler](https://github.com/google/protobuf)
- [protoc-gen-go plugin](https://github.com/golang/protobuf)
- [docker for mac](https://store.docker.com/editions/community/docker-ce-desktop-mac)
- [docker for mac Kubernetes cluster](https://docs.docker.com/docker-for-mac/#kubernetes) minikube should also work, note that I used the edge version of docker for mac
- [helm](https://github.com/helm/helm#install)
- [istio](https://istio.io/docs/setup/kubernetes/helm-install/) note that I included the required Istio 1.0.4 helm files in the charts folder for simplicity

# How to build it

## Locally

``` shell
go get github.com/EtienneDufresne/grpc-k8s-lb
make generate-protos
make build
```

## Docker

``` shell
make build-docker
```

# How to run it

## Go binary

Start the server:
``` shell
grpc-k8s-lb -s
```

Start the client
``` shell
grpc-k8s-lb
```

## Helm chart

This assumes you have a clean Docker for Mac / Minikube Kubernetes cluster and that you built the echo client/server docker image.

Start two gRPC echo servers:
``` shell
helm upgrade --install grpc-server ./chart \
  --set replicaCount=2
```

Start a gRPC echo client:
``` shell
helm upgrade --install grpc-client ./chart \
  --set args[0]=-h,args[1]=grpc-server-grpc-k8s-lb:8080
```

At this point you should see that every 5 seconds, the client is making a request to the same grpc server pod as there is no load balancing:
``` shell
kubectl logs -f grpc-client-xxx
```

Delete all helm releases from the default namespace
``` shell
helm delete --purge $(helm list -q)
```

Setup the service account that will be used by the helm tiller to setup istio:
``` shell
kubectl apply -f charts/helm-service-account.yaml
```

Setup helm tiller to use the service account:
``` shell
helm init --service-account tiller
```

Install istio in the istio-system namespace:
``` shell
helm install charts/istio --name istio --namespace istio-system --set tracing.enabled=true
```

Turn on auto Envoy proxy sidecar injection on the default namespace:
``` shell
kubectl label namespace default istio-injection=enabled
```

Start two gRPC echo servers:
``` shell
helm upgrade --install grpc-server charts/echo \
  --set replicaCount=2
```

Start a gRPC echo client:
``` shell
helm upgrade --install grpc-client charts/echo \
  --set args[0]=-h,args[1]=grpc-server-grpc-k8s-lb:8080
```

At this point you should see that every 5 seconds, the client is making a request to different grpc server pod as istio/envoy load balancing is enabled:
``` shell
kubectl logs -f grpc-client-xxx
```

## Clean Up

Delete all helm releases from the default namespace
``` shell
helm delete --purge $(helm list -q)
```

Delete Istio Custom Resource Definitions
``` shell
kubectl delete -f istio/install/kubernetes/helm/istio/templates/ -n istio-system
```

## TODO

- Add a simple [grcp-web](https://github.com/grpc/grpc-web) UI
- Investigate request tracing and service mesh observability with Jaeger and Kialia

## References

- [Intro to protobuf and gRPC](https://medium.com/@shijuvar/building-high-performance-apis-in-go-using-grpc-and-protocol-buffers-2eda5b80771b)
- [gRCP chat implemented in go](https://github.com/rodaine/grpc-chat)
- [gRCP Load Balancing](https://kubernetes.io/blog/2018/11/07/grpc-load-balancing-on-kubernetes-without-tears/)
- [Istio service mesh with Envoy as gRPC proxy example](https://istio.io/docs/examples/bookinfo/)
- [Istio & Envoy grpc metrics](https://medium.com/pismolabs/istio-envoy-grpc-metrics-winning-with-service-mesh-in-practice-d67a08acd8f7)
- [HTTP2 vs Webscocket](https://www.infoq.com/articles/websocket-and-http2-coexist)
- [grcp-web](https://github.com/grpc/grpc-web)
- [IETF gRCP RFC](https://tools.ietf.org/html/draft-kumar-rtgwg-grpc-protocol-00)
- [IETF websocket RFC](https://tools.ietf.org/html/rfc6455)
