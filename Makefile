
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	script/gen.sh

.PHONY: build
build: proto

	go build -o nebula main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t nebula:latest

.PHONY: run
run:
	docker run --name nebula -d -p 8080:8080 -v /Users/zuoyi-macpro/Data/nebula/runtime:/runtime nebula
