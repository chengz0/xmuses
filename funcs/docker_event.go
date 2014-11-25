package funcs

import (
	"encoding/json"
	"github.com/deepglint/go-dockerclient"
	"log"
	"sync"
	"time"
)

var (
	ContainersIN map[string]string
	ContainersIS map[string]string
	Listener     chan *docker.APIEvents
	timeout      = time.After(1 * time.Second)
	StatusSync   *sync.Mutex
)

func InitContainerEvent(client *docker.Client) {
	// add listener
	err := client.AddEventListener(Listener)
	if err != nil {
		log.Fatalf("Error adding docker listener: %s", err.Error())
	}
	defer func() {
		err = client.RemoveEventListener(Listener)
		if err != nil {
			log.Fatalf("Error removing docker listener: %s", err.Error())
		}
	}()

	// listening events
	for {
		select {
		case msg := <-Listener:
			// log.Println(msg)
			body, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Can not parse apievent: %s", err.Error())
				break
			}
			log.Println(string(body))

			StatusSync.Lock()
			ContainersIS[msg.ID] = msg.Status
			StatusSync.Unlock()
		case <-timeout:
			break
		}
	}

}
