package main

import (
	"fmt"
	"github.com/Wall-js/nebula"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/rabbitmq"

	"log"
	"time"
)

var topic = "nebula"

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := nebula.RabbitBroker.Publish(topic, msg); err != nil {
			log.Printf("[pub] 发布消息失败： %v", err)
		} else {
			fmt.Println("[pub] 发布消息：", string(msg.Body))
		}
		i++
	}
}

func sub() {
	_, err := nebula.RabbitBroker.Subscribe(topic, func(p broker.Event) error {
		//_, err := nebula.RabbitBroker.Subscribe("sample", func(p broker.Event) error {
		fmt.Printf("[sub] 订阅 Body: %s, Header: %s", string(p.Message().Body), p.Message().Header)
		return nil
	},
		broker.Queue("mu.micro.book.topic.queue"),
	)
	if err != nil {
		fmt.Println(err)
	}
}

func day() {
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		nebula.Publish(nebula.RabbitBroker, "sample", []byte("test..."))
		i++
	}
}

func testSub(e broker.Event) error {
	fmt.Printf("[sub] 订阅 Body: %s, Header: %s \n", string(e.Message().Body), e.Message().Header)
	return nil
}

func testPub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		//nebula.Publish(nebula.RabbitBroker, topic, fmt.Sprintf("%s", time.Now().String()))
		nebula.Publish(nebula.RabbitBroker, topic, []byte("supplierId:66"))
		i++
	}
}

func main() {
	nebula.InitRabbitBroker()

	//nebula.Subscribe(nebula.RabbitBroker, topic, testSub)
	nebula.Subscribe(nebula.RabbitBroker, topic, testSub, broker.Queue("mu.micro.book.topic.queue"))
	go testPub()

	nebula.Run()
}

//<-time.After(time.Second * 10)
//	select {}
//}
//func main() {
//
//	if err := broker.Init(); err != nil {
//		log.Fatalf("Broker 初始化错误：%v", err)
//	}
//	if err := broker.Connect(); err != nil {
//		log.Fatalf("Broker 连接错误：%v", err)
//	}
//
//	go pub()
//go sub()
//<-time.After(time.Second * 10)
//	select {}
//}
