module github.com/Wall-js/nebula

go 1.12

replace (
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
	github.com/nats-io/nats.go => github.com/nats-io/nats.go v1.8.1
)

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/golang/protobuf v1.3.2
	github.com/lucas-clemente/quic-go v0.11.2 // indirect
	github.com/micro/go-micro v1.8.1
	github.com/micro/go-plugins v1.1.1
	github.com/satori/go.uuid v1.2.0
	google.golang.org/grpc v1.22.0
)
