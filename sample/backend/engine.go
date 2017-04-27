package backend

import (
	"log"
	"net/http"

	"github.com/intercom/speakeasy"
)

type SampleEngine struct {
	AppState     *AppState
	RequestCount int
	AlwaysFail   bool
}

func NewEngine() *SampleEngine {
	engine := &SampleEngine{}
	engine.Reset()
	return engine
}

func (engine *SampleEngine) Reset() {
	engine.RequestCount = 0
	engine.AppState = DefaultAppState()
}

func (engine *SampleEngine) WrapHandler(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + " " + r.URL.EscapedPath())
		engine.RequestCount++
		if engine.AlwaysFail {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			f(w, r)
		}
	}
}

func (engine *SampleEngine) Setup(server *speakeasy.Server) {
	server.Router.HandleFunc("/channels", engine.WrapHandler(engine.channelsResponse)).Methods("GET")
	server.Router.HandleFunc("/channels", engine.WrapHandler(engine.newChannelResponse)).Methods("POST")
	server.Router.HandleFunc("/channels/{channelName}", engine.WrapHandler(engine.channelByNameResponse)).Methods("GET")
	server.Router.HandleFunc("/channels/{channelName}/reply", engine.WrapHandler(engine.replyResponse)).Methods("POST")
	server.Router.HandleFunc("/assets/{assetDirectory}/{assetFilename}", engine.WrapHandler(engine.assetResponse)).Methods("GET")
}
