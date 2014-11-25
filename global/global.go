package global

import (
	// "os"
	"github.com/chengz0/xmuses/funcs"
	"github.com/deepglint/go-dockerclient"
	"sync"
)

// var host = os.Getenv("IP")
var (
	Host string
)

func InitGlobal() {
	Host = "172.16.110.134"

	funcs.ContainersMap = make(map[string]funcs.ContainerStateStruct)

	funcs.ContainersIN = make(map[string]string)
	funcs.ContainersIS = make(map[string]string)
	funcs.StatusSync = new(sync.Mutex)

	funcs.Listener = make(chan *docker.APIEvents)
}
