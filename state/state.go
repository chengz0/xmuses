package state

import (
	"github.com/chengz0/xmuses/global"
	"github.com/deepglint/go-dockerclient"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"log"
	"net/http"
)

type ContainerStateStruct struct {
	Id       string
	Image    string
	Cmd      []string
	Running  bool
	ExitCode int
}

var (
	martini_m     *martini.ClassicMartini
	client        *docker.Client
	ContainersMap map[string]ContainerStateStruct
)

func StartMartini() {
	// martini
	martini_m = martini.Classic()
	martini_m.Use(render.Renderer(render.Options{
		Funcs: []template.FuncMap{{
			"nl2br":      nl2br,
			"htmlquote":  htmlQuote,
			"str2html":   str2html,
			"dateformat": dateFormat,
		}},
	}))

	// system state
	martini_m.Get("/system", func() bool {
		return true
	})

	// docker client
	ContainersMap = make(map[string]ContainerStateStruct)
	client, err := docker.NewClient("http://" + global.Host + ":4243")
	if err != nil {
		log.Printf("Error creating client: %s", err.Error())
		return
	}
	log.Println(client.Version())

	Containers(client)

	// docker container state
	martini_m.Get("/container/:name", func(params martini.Params) bool {
		name := params["name"]
		state := ContainerState(client, name)
		log.Println(state)
		return state
	})

	martini_m.Get("/containers", func() map[string]bool {
		ret := make(map[string]bool)
		for k, v := range ContainersMap {
			ret[k] = v.Running
		}
		log.Println(ret)
		return ret
	})

	http.ListenAndServe(":3000", martini_m)
}
