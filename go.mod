module github.com/Wall-js/nebula

go 1.12

replace (
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
	github.com/nats-io/nats.go => github.com/nats-io/nats.go v1.8.1
)

require (
	cloud.google.com/go v0.41.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-kit/kit v0.9.0 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.17.1
	github.com/micro/go-plugins v1.1.1
	github.com/satori/go.uuid v1.2.0
	golang.org/x/exp v0.0.0-20190627132806-fd42eb6b336f // indirect
	golang.org/x/image v0.0.0-20190703141733-d6a02ce849c9 // indirect
	golang.org/x/mobile v0.0.0-20190711165009-e47acb2ca7f9 // indirect
	google.golang.org/grpc v1.25.1
)
