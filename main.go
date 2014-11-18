package main

import (
	"./publish"
	"github.com/chengz0/xmuses/global"
	"github.com/chengz0/xmuses/state"
	"github.com/deepglint/go-nsq"
	"log"
)

func main() {

	// config
	global.InitGlobal()
	log.Println(global.Host)

	// mamrtini
	state.StartMartini()

	mq_cfg := nsq.NewConfig()
	mq_cfg.ClientID = "client_id"

	nsq_url := global.Host + ":4150"
	producer, err := nsq.NewProducer(nsq_url, mq_cfg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer producer.Stop()

	sys := &publish.SystemState{
		SensorId: "cq.dsj.s1",
		Hostname: "10.167.13.152",
		State:    true,
	}
	log.Println(sys)

	topic := "topic"
	publish.PubSystem(producer, topic, sys)
}
