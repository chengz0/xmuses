package state

import (
	"github.com/deepglint/go-dockerclient"
	"log"
	"strings"
)

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
			log.Printf("Error inspecting containers: %s", err.Error())
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
	}
}

func ContainerState(client *docker.Client, name string) bool {
	container, err := client.InspectContainer(ContainersMap[name].Id)
	if err != nil {
		log.Printf("Error inspecting containers: %s", err.Error())
		return false
	}
	return container.State.Running
}
