package nebula

import (
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/broker/nats"
	"github.com/micro/go-plugins/broker/nsq"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/satori/go.uuid"
)

var (
	RabbitBroker = rabbitmq.NewBroker()
	NatsBroker   = nats.NewBroker()
	NsqBroker    = nsq.NewBroker()
)

func NewRabbitBroker(opts ...broker.Option) broker.Broker {
	return rabbitmq.NewBroker(opts...)
}

func NewNatsBroker(opts ...broker.Option) broker.Broker {
	return nats.NewBroker(opts...)
}

func NewNsqBroker(opts ...broker.Option) broker.Broker {
	return nsq.NewBroker(opts...)
}

func InitRabbitBroker(b broker.Broker, options ...broker.Option) {
	rabbitBrokerAddr := Conf.Get("rabbitBrokerAddr").String("")
	options = append(options, broker.Addrs(rabbitBrokerAddr))
	if err := b.Init(
		options...,
	); err != nil {
		log.Fatalf("Broker 初始化错误：%v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker 连接错误：%v", err)
	}
}

func InitNatsBroker(b broker.Broker, options ...broker.Option) {
	natsBrokerAddr := Conf.Get("natsBrokerAddr").String("")
	options = append(options, broker.Addrs(natsBrokerAddr))
	if err := b.Init(
		options...,
	); err != nil {
		log.Fatalf("Broker 初始化错误：%v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker 连接错误：%v", err)
	}
}

func InitNsqBroker(b broker.Broker, options ...broker.Option) {
	nsqBrokerAddr := Conf.Get("nsqBrokerAddr").String("")
	options = append(options, broker.Addrs(nsqBrokerAddr))
	if err := b.Init(
		options...,
	); err != nil {
		log.Fatalf("Broker 初始化错误：%v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker 连接错误：%v", err)
	}
}

func Publish(b broker.Broker, topic string, data []byte, opts ...broker.PublishOption) string {

	id := uuid.NewV4().String()
	msg := &broker.Message{
		Header: map[string]string{
			"id": id,
		},
		Body: data,
	}

	if err := b.Publish(topic, msg, opts...); err != nil {
		log.Logf("[pub] 发布消息失败： %v", err)
	} else {
		log.Log("[pub] 发布消息：", string(msg.Body))
	}

	return id
}

func subscribe(b broker.Broker, topic string, handle func(broker.Event) error, opts ...broker.SubscribeOption) {
	_, err := b.Subscribe(
		topic,
		handle,
		opts...,
	)
	if err != nil {
		log.Log(err)
	}
}

func Subscribe(b broker.Broker, topic string, handle func(broker.Event) error, opts ...broker.SubscribeOption) {
	go subscribe(b, topic, handle, opts...)
}
