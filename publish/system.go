package publish

import (
	"encoding/json"
	"github.com/deepglint/go-nsq"
	"log"
	"reflect"
)

type SystemState struct {
	SensorId string
	Hostname string
	State    bool
}

func PubSystem(producer *nsq.Producer, topic string, state interface{}) {

	log.Println(reflect.TypeOf(state))

	body, err := json.Marshal(state)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(string(body))

	// producer.Publish(topic, string(body))
}
