package state

import (
	"github.com/chengz0/xmuses/funcs"
	// "github.com/martini-contrib/render"
	"log"
	"net/http"
)

func EventRouter() {
	martini_m.Get("/container/events", func(w http.ResponseWriter, req *http.Request) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		ns := make(map[string]string)
		for k, v := range funcs.ContainersIS {
			ns[funcs.ContainersIN[k]] = v
		}
		body := RenderDataDto(ns)
		funcs.StatusSync.Lock()
		for k, _ := range funcs.ContainersIS {
			delete(funcs.ContainersIS, k)
		}
		funcs.StatusSync.Unlock()

		return body
	})
}
