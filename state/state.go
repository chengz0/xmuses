package state

import (
	"encoding/json"
	"github.com/chengz0/xmuses/funcs"
	"github.com/chengz0/xmuses/global"
	"github.com/deepglint/go-dockerclient"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"log"
	"net/http"
)

var (
	martini_m *martini.ClassicMartini
	client    *docker.Client
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
	martini_m.Get("/system", func() string {
		return "true"
	})

	// docker client
	client, err := docker.NewClient("http://" + global.Host + ":4243")
	if err != nil {
		log.Printf("Error creating client: %s", err.Error())
		return
	}
	log.Println(client.Version())

	funcs.Containers(client)

	go funcs.InitContainerEvent(client)

	// // docker container state
	// martini_m.Get("/container/:name", func(params martini.Params) string {
	// 	name := params["name"]
	// 	state := funcs.ContainerState(client, name)
	// 	if state {
	// 		return "true"
	// 	} else {
	// 		return "false"
	// 	}
	// })

	martini_m.Get("/containers", func() string {
		ret, err := json.Marshal(funcs.ContainersMap)
		if err != nil {
			log.Printf("Error marshalling containers: %s", err.Error())
			return ""
		}
		return string(ret)
	})

	EventRouter()

	http.ListenAndServe(":3000", martini_m)
}

type Dto struct {
	Succ bool
	Msg  string
	Data interface{}
}

func ErrDto(message string) Dto {
	return Dto{Succ: false, Msg: message}
}

func DataDto(d interface{}) Dto {
	return Dto{Succ: true, Msg: "", Data: d}
}

func RenderErrDto(message string) string {
	dto := ErrDto(message)
	bs, err := json.Marshal(dto)
	if err != nil {
		return err.Error()
	} else {
		return string(bs)
	}
}

func RenderDataDto(d interface{}) string {
	dto := DataDto(d)
	bs, err := json.Marshal(dto)

	if err != nil {
		return err.Error()
	} else {
		return string(bs)
	}
}
