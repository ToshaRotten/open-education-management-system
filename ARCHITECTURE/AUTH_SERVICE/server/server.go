package server

import (
	"encoding/json"
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/buffer"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/models"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/config"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	Config *config.Config
	Logger *logrus.Logger
	Buffer *buffer.Buffer
	Router *mux.Router
	Users  *UserManager.UserManager
}

func New() *APIServer {
	var s APIServer
	s.Config = config.New()
	s.Logger = logrus.New()
	s.Router = mux.NewRouter()
	s.Buffer = buffer.New()
	s.Users = UserManager.New()
	return &s
}

func (s *APIServer) Start(config *config.Config) error {
	s.Config = config
	err := s.configureLogger()
	if err != nil {
		s.Logger.Error(err)
		return err
	}
	s.configureRouter()
	s.configureFileHelper()
	s.Logger.Info("Server is started ...")
	s.Logger.Info("Bind addr: http://", s.Config.Host+s.Config.Port)
	err = http.ListenAndServe(s.Config.Host+s.Config.Port, s.Router)
	if err != nil {
		s.Logger.Error(err)
		return err
	}
	return nil
}

func (s *APIServer) configureLogger() error {
	s.Logger.Trace("Logger ...")
	s.Logger.SetLevel(logrus.Level(6))
	return nil
}

func (s *APIServer) configureFileHelper() {

}

// configureRouter ..
func (s *APIServer) configureRouter() {
	s.Logger.Trace("Router ...")
	s.Router.HandleFunc("/user/create", s.create())
	s.Router.HandleFunc("/user/read", s.read())
	s.Router.HandleFunc("/user/update", s.update())
	s.Router.HandleFunc("/user/delete", s.delete())
}

func (s *APIServer) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usersData, err := utils.ParseUsersFromJSON(r)
		if err != nil {
			s.Logger.Error(err)
		}
		s.Users.CreateUsers(usersData.Users)
	})
}

func (s *APIServer) read() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usersData, err := utils.ParseUsersFromJSON(r)
		if err != nil {
			s.Logger.Error(err)
		}
		users := s.Users.ReadUsers(usersData.Users)
		response, err := json.Marshal(users)
		if err != nil {
			s.Logger.Error(err)
		}
		_, err = w.Write(response)
		if err != nil {
			s.Logger.Error(err)
		}
	})
}

func (s *APIServer) update() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var update struct {
			OldUserData []models.User `json:"oldUsersData"`
			NewUserData []models.User `json:"newUsersData"`
		}
		data, err := io.ReadAll(r.Body)

		fmt.Println(string(data))
		if err != nil {
			s.Logger.Error(err)
		}
		err = json.Unmarshal(data, &update)
		if err != nil {
			s.Logger.Error(err)
		}
		s.Users.UpdateUsers(update.OldUserData, update.NewUserData)
	})
}

func (s *APIServer) delete() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usersData, err := utils.ParseUsersFromJSON(r)
		if err != nil {
			s.Logger.Error(err)
		}
		s.Users.DeleteUsers(usersData.Users)
	})
}
