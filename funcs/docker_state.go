package funcs

import (
	"github.com/deepglint/go-dockerclient"
	"log"
	"strings"
)

var ContainersMap map[string]ContainerStateStruct

type ContainerStateStruct struct {
	Id       string
	Image    string
	Cmd      []string
	Running  bool
	ExitCode int
}

func Containers(client *docker.Client) {
	//
	cs, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		log.Printf("Error listing containers: %s", err.Error())
		return
	}
	for _, c := range cs {
		container, err := client.InspectContainer(c.ID)
		if err != nil {
			log.Printf("Error inspecting container [Containers]: %s", err.Error())
			return
		}
		css := new(ContainerStateStruct)
		css.Id = container.ID
		css.Image = container.Image
		css.Cmd = container.Config.Cmd
		css.Running = container.State.Running
		css.ExitCode = container.State.ExitCode

		name := strings.Replace(container.Name, "/", "", -1)
		ContainersMap[name] = *css
		//
		ContainersIN[container.ID] = name
	}
}

func ContainerState(client *docker.Client, name string) bool {
	if v, ok := ContainersMap[name]; ok {
		container, err := client.InspectContainer(v.Id)
		if err != nil {
			log.Printf("Error inspecting container [ContainerState]: %s", err.Error())
			return false
		}
		return container.State.Running
	}
	return false
}
