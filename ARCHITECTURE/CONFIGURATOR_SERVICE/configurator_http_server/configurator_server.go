package configurator_http_server

import (
	"fmt"
	"github.com/gorilla/mux"
	"main/configurator"
	"net/http"
)

type Server struct {
	Configurator *configurator.Configurator
	Router       *mux.Router
}

func New() *Server {
	var s Server
	s.Router = mux.NewRouter()
	s.Configurator = configurator.New()
	return &s
}

func (s *Server) Start() {
	port := SearchAvailablePort()
	fmt.Println("Server is started")
	fmt.Println("Server bind addr: http://" + "localhost:" + port)
	err := http.ListenAndServe("localhost:"+port, s.Router)
	if err != nil {
		fmt.Println("ListenAndServe Error: ", err)
	}
}

// configureRouter ..
func (s *Server) configureRouter() {
	s.Router.HandleFunc("/user/create", s.getConfigs())
}

func (s *Server) getConfigs() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
