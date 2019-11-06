package nebula

import (
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/broker/nats"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/satori/go.uuid"
)

var RabbitBroker = rabbitmq.NewBroker()
var NatsBroker = nats.NewBroker()

func InitRabbitBroker(options ...broker.Option) {
	rabbitBrokerAddr := Conf.Get("rabbitBrokerAddr").String("")
	options = append(options, broker.Addrs(rabbitBrokerAddr))
	if err := RabbitBroker.Init(
		options...,
	); err != nil {
		log.Fatalf("Broker 初始化错误：%v", err)
	}
	if err := RabbitBroker.Connect(); err != nil {
		log.Fatalf("Broker 连接错误：%v", err)
	}

}

func InitNatsBroker(options ...broker.Option) {
	rabbitBrokerAddr := Conf.Get("natsBrokerAddr").String("")
	options = append(options, broker.Addrs(rabbitBrokerAddr))
	if err := NatsBroker.Init(
		options...,
	); err != nil {
		log.Fatalf("Broker 初始化错误：%v", err)
	}
	if err := RabbitBroker.Connect(); err != nil {
		log.Fatalf("Broker 连接错误：%v", err)
	}

}

func Publish(b broker.Broker, topic string, data string) string {

	id := uuid.NewV4().String()
	msg := &broker.Message{
		Header: map[string]string{
			"id": id,
		},
		Body: []byte(data),
	}

	if err := b.Publish(topic, msg); err != nil {
		log.Logf("[pub] 发布消息失败： %v", err)
	} else {
		log.Log("[pub] 发布消息：", string(msg.Body))
	}

	return id
}
