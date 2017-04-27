package speakeasy

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Engine Engine
	Router *mux.Router
}

func NewServer(engine Engine) *Server {
	return &Server{
		Engine: engine,
		Router: mux.NewRouter().StrictSlash(true),
	}
}

type Engine interface {
	Setup(server *Server)
}

func (server *Server) Start() {
	server.Engine.Setup(server)
	log.Fatal(http.ListenAndServe(":3000", server.Router))
}
