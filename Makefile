.PHONY: build
build:
	go build .

.PHONY: generate-protos
generate-protos:
	protoc --go_out="plugins=grpc:." protos/message.proto

.PHONY: build-docker
build-docker:
	docker build --rm -t etiennedufresne/grpc-k8s-lb .
