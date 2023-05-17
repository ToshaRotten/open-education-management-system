package statistic_http_server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"main/statistic_http_server/models"
	"main/statistic_http_server/telegram"
	"net/http"
)

type Server struct {
	telegram *telegram.Telegram
	Router   *mux.Router
}

func New() *Server {
	var s Server
	s.telegram = telegram.New()
	s.Router = mux.NewRouter()
	return &s
}

func (s *Server) Start() {
	s.configureRouter()
	err := http.ListenAndServe(":9908", s.Router)
	if err != nil {
		fmt.Println("ListenAndServe error: ", err)
	}
}

func (s *Server) configureRouter() {
	s.Router.HandleFunc("/log/send", s.send())
}

func (s *Server) send() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var temp models.Log
		data, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error read: ", err)
		}
		err = json.Unmarshal(data, &temp)
		if err != nil {
			fmt.Println("Unmarshal error: ", err)
		}
		s.telegram.SendMessage(temp.Message)
	})
}
