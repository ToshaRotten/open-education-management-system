package auth_http_server

import (
	"encoding/json"
	"errors"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server/UserManager"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server/UserManager/buffer"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server/UserManager/models"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/auth_http_server/utils"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/configurator_http_client"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/statistic_http_client"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

type APIServer struct {
	TGLog  *statistic_http_client.StatisticClient
	Config *configurator_http_client.Config
	Logger *logrus.Logger
	Buffer *buffer.Buffer
	Router *mux.Router
	Users  *UserManager.UserManager
}

func New() *APIServer {
	var s APIServer
	s.TGLog = statistic_http_client.New()
	s.Config = configurator_http_client.NewConfig()
	s.Logger = logrus.New()
	s.Router = mux.NewRouter()
	s.Buffer = buffer.New()
	s.Users = UserManager.New()
	return &s
}

func (s *APIServer) Start(config *configurator_http_client.Config) error {
	s.Config = config
	err := s.configureLogger()
	if err != nil {
		s.Logger.Error(err)
		return err
	}
	s.configureUserManager()
	s.configureRouter()
	s.configureFileHelper()
	s.Logger.Info("Server is started ...")
	s.Logger.Info("Server bind addr: http://185.21.142.92:8079")
	s.TGLog.SendLog("Server is started")
	err = http.ListenAndServe("185.21.142.92:8079", s.Router)
	if err != nil {
		return err
	}
	return nil
}

func (s *APIServer) configureUserManager() {
	s.Users.SetScheme("auth_http_server/UserManager/database/test.db")
	s.Users.InitManager()
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

	s.Router.HandleFunc("/user/register", s.register())
	s.Router.HandleFunc("/user/auth", s.auth())
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
		s.TGLog.SendLog("Create request")
		usersData, err := utils.ParseUsersFromJSON(r)
		if err != nil {
			s.Logger.Error(err)
		}
		err = s.Users.CreateUsers(usersData.Users)
		if err != nil {
			s.Logger.Error(err)
			w.WriteHeader(http.StatusNotAcceptable)
		}
		w.WriteHeader(http.StatusCreated)
		s.TGLog.SendLog("User is created")
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
		s.TGLog.SendLog("Read request, written data to client:" + string(response))
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

		if err != nil {
			s.Logger.Error(err)
		}
		err = json.Unmarshal(data, &update)
		if err != nil {
			s.Logger.Error(err)
		}
		s.TGLog.SendLog("Update request, data:" + string(data))
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
		requestData, _ := io.ReadAll(r.Body)
		s.TGLog.SendLog("Delete user request. request data: " + string(requestData))
		usersData, err := utils.ParseUsersFromJSON(r)
		if err != nil {
			s.Logger.Error(err)
		}
		s.Users.DeleteUsers(usersData.Users)
		s.TGLog.SendLog("User is deleted")
	})
}

func (s *APIServer) auth() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(204)
			return
		}
		requestData, _ := io.ReadAll(r.Body)
		s.TGLog.SendLog("Auth request, request data: " + string(requestData))
		userData, err := utils.ParseUsersFromJSON(r)
		if err != nil {
			s.Logger.Error(err)
		}
		valid := utils.ValidateUsers(userData.Users)
		if valid {
			foundedUsers := s.Users.ReadUsers(userData.Users)
			if len(foundedUsers) > 1 {
				s.Logger.Error(errors.New("Error auth"))
				s.TGLog.SendLog("Auth request: wrong input")
			} else {
				if (foundedUsers[0].Hash == userData.Users[0].Hash) &&
					(foundedUsers[0].Email == userData.Users[0].Email) {
					marshaled, _ := json.Marshal(&foundedUsers[0])
					w.Write(marshaled)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
				}
			}
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			s.TGLog.SendLog("Auth request: user invalid")
		}
	})
}

func (s *APIServer) register() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(204)
			return
		}
		requestData, _ := io.ReadAll(r.Body)
		s.TGLog.SendLog("Register request, request data: " + string(requestData))
		userData, err := utils.ParseUsersFromJSON(r)
		if err != nil {
			s.Logger.Error(err)
		}
		valid := utils.ValidateUsers(userData.Users)
		if valid {
			err = s.Users.CreateUsers(userData.Users)
			if err != nil {
				s.Logger.Error(err)
				w.WriteHeader(http.StatusNotAcceptable)
			}
			w.WriteHeader(http.StatusCreated)
			s.TGLog.SendLog("Register request: user created")
		}
	})
}
