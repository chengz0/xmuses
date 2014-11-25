package main

import (
	"encoding/json"
	"github.com/chengz0/xmuses/funcs"
	"github.com/gopherjs/jquery"
	"log"
	"strings"
	"time"
)

var (
	timeout = time.After(1 * time.Second)
	jQuery  = jquery.NewJQuery
)

const (
	TEST          = "li#test"
	UBUNTU        = "li#ubuntu"
	LIBRA         = "li#libra"
	WEEDMASTER    = "li#weedmaster"
	WEEDVOLUME1   = "li#weedvolume1"
	REDIS         = "li#redis"
	NSQD          = "li#nsqd"
	SENSORCLOUD   = "li#sensorcloud"
	EVENTUPLOAD   = "li#eventupload"
	SLICESERVER   = "li#sliceserver"
	SLICEUPLOAD   = "li#sliceupload"
	DOCKERMANAGER = "li#dockermanager"
	DOCKERMONITOR = "li#dockermonitor"
	NSQCLASSIFIER = "li#nsqclassifier"
	INFLUXDB      = "li#influxdb"
	CADVISOR      = "li#cadvisor"
	STATS         = "li#stats"
	ELASTICSEARCH = "li#elasticsearch"
	KIBANA3       = "li#kibana3"
	FLUENTD       = "li#fluentd"
	RAWETCD       = "li#rawetcd"
	DBETCD        = "li#dgetcd"
	NTPSYNC       = "li#ntpsync"
	RUNTIMESERVER = "li#runtimeserver"
	SCHEDULER     = "li#scheduler"
)

func main() {
	// jquery
	log.Println(jQuery().Jquery)

	// listening events
	for {
		select {
		case msg := <-funcs.Listener:
			// log.Println(msg)
			body, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Can not parse apievent: %s", err.Error())
				continue
			}
			log.Println(string(body))
			curcontainer := strings.ToUpper(funcs.ContainersIN[msg.ID])
			jQuery(curcontainer).ToggleClass("running")
		case <-timeout:
			break
		}
	}
}
