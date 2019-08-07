
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	./proto/gen.sh

.PHONY: build
build: proto

	go build -o nebula main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t nebula:latest
