package configurator_http_server

import (
	"encoding/json"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/CONFIGURATOR_SERVICE/configurator"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/CONFIGURATOR_SERVICE/configurator/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Server struct {
	Configurator *configurator.Configurator
	Router       *mux.Router
	Log          *logrus.Logger
}

func New() *Server {
	var s Server
	s.Router = mux.NewRouter()
	s.Configurator = configurator.New("configurator/base_config/default.yaml")
	s.Log = logrus.New()
	return &s
}

func (s *Server) Start() {
	s.configureRouter()
	port := SearchAvailablePort()
	s.Log.Info("Server is started")
	s.Log.Info("Server bind addr: http://" + "localhost:" + port)
	err := http.ListenAndServe("localhost:"+port, s.Router)
	if err != nil {
		s.Log.Error("ListenAndServe Error: ", err)
	}
}

// configureRouter ..
func (s *Server) configureRouter() {
	s.Router.HandleFunc("/config", s.getConfig())
	s.Router.HandleFunc("/health", s.health())
}

func (s *Server) health() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("it works"))
		if err != nil {
			s.Log.Error(err)
		}
	})
}

func (s *Server) getConfig() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var conf models.Config
		data, err := io.ReadAll(r.Body)
		if err != nil {
			s.Log.Error(err)
		}
		err = json.Unmarshal(data, &conf)
		if err != nil {
			s.Log.Error(err)
		}

		s.Log.Info("get config request, request data:", string(data))

		respConfig := s.Configurator.Buffer.GetService(conf.ServiceName)
		data, err = json.Marshal(respConfig)
		if err != nil {
			s.Log.Error(err)
		}
		_, err = w.Write(data)
		if err != nil {
			s.Log.Error(err)
		}
	})
}
