module github.com/Wall-js/nebula

go 1.12

replace (
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
	github.com/nats-io/nats.go => github.com/nats-io/nats.go v1.8.1
)

require (
	github.com/micro/go-micro v1.7.1-0.20190627135301-d8e998ad85fe
	github.com/micro/go-plugins v1.1.1
)
