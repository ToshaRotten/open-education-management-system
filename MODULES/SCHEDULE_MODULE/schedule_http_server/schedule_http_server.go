package schedule_http_server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"main/schedule_http_server/models"
	"main/schedule_http_server/schedule_controller"
	"net/http"
)

type Server struct {
	scheduler *schedule_controller.Controller
	Router    *mux.Router
}

func New() *Server {
	var s Server
	s.scheduler = schedule_controller.New()
	s.Router = mux.NewRouter()
	return &s
}

func (s *Server) configureRouter() {
	s.Router.HandleFunc("/schedule/create", s.create())
	s.Router.HandleFunc("/schedule/read", s.read())
}

func (s *Server) Start() {
	s.configureRouter()
	err := http.ListenAndServe(":9909", s.Router)
	if err != nil {
		fmt.Println("ListenAndServe error: ", err)
	}
}

func (s *Server) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var temp []models.Lesson
		data, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(data, &temp)
		if err != nil {
			fmt.Println(err)
		}
		s.scheduler.Create(data)
	})
}

func (s *Server) read() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := s.scheduler.Read()
		_, err := w.Write(data)
		if err != nil {
			fmt.Println(err)
		}
	})
}
