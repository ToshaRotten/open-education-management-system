package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/buffer"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/UserManager/models"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/config"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/server/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
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
	s.Logger = &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			ForceColors: true,
			ForceQuote:  true,
		},
	}

	s.Logger.SetLevel(logrus.Level(6))
	s.Logger.Trace("Logger ...")
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(204)
			return
		}
		usersData, err := utils.ParseUsersFromJSON(r)
		if err != nil {
			s.Logger.Error(err)
		}
		s.Users.CreateUsers(usersData.Users)
	})
}

func (s *APIServer) read() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(204)
			return
		}
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(204)
			return
		}
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(204)
			return
		}
		usersData, err := utils.ParseUsersFromJSON(r)
		if err != nil {
			s.Logger.Error(err)
		}
		s.Users.DeleteUsers(usersData.Users)
	})
}

func (s *APIServer) auth() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userData, err := utils.ParseUsersFromJSON(r)
		if err != nil {
			s.Logger.Error(err)
		}
		valid := s.Users.ValidateUsers(userData.Users)
		if valid {
			foundedUsers := s.Users.ReadUsers(userData.Users)
			if len(foundedUsers) > 1 {
				s.Logger.Error(errors.New("Хана"))
			} else {
				if foundedUsers[0].Hash == userData.Users[0].Hash {
					marshaled, _ := json.Marshal(&userData.Users[0])
					w.Write(marshaled)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
				}
			}
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
		}
	})
}

func (s *APIServer) register() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//userData, err := utils.ParseUsersFromJSON(r)
		//if err != nil {
		//	s.Logger.Error(err)
		//}
	})
}
