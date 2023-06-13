package root_http_server

import (
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/ROOT_SERVICE/configurator_http_client"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	Router             *mux.Router
	Log                *logrus.Logger
	ConfiguratorClient *configurator_http_client.Client
}

func New() *Server {
	var s Server
	s.ConfiguratorClient = configurator_http_client.New()
	s.Router = mux.NewRouter()
	s.Log = logrus.New()
	return &s
}

func (s *Server) Start() {
	config := s.ConfiguratorClient.GetConfig()
	s.Log.Info("Server is started")
	s.Log.Info("Server bind addr: http://" + config.Host + ":" + config.Port)
	err := http.ListenAndServe(config.Host+":"+config.Port, s.Router)
	if err != nil {
		s.Log.Error("ListenAndServe Error: ", err)
	}
}

// configureRouter ..
func (s *Server) configureRouter() {
	s.Router.HandleFunc("/runner", s.runner())
}

func (s *Server) runner() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
