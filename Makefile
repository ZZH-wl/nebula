
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	script/gen.sh

.PHONY: build
build: proto

	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -i -o nebula main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t nebula:latest

.PHONY: run
run:
	docker run --name nebula -d -p 8080:8080 -v /Users/zuoyi-macpro/Data/nebula/runtime:/runtime nebula
