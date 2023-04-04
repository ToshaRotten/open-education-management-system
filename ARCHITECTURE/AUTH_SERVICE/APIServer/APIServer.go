package APIServer

import (
	"encoding/json"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/UserManager"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/UserManager/buffer"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/UserManager/models"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/config"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
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
	s.Router.HandleFunc("/user/register", s.register())
	s.Router.HandleFunc("/user/auth", s.auth())
	s.Router.HandleFunc("/user/edit", s.edit())
	s.Router.HandleFunc("/user/list", s.userList())
}

func (s *APIServer) userList() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		array := s.Users.GetUserList()
		data, err := json.Marshal(&array)
		if err != nil {
			s.Logger.Error(err)
		}
		_, err = w.Write(data)
		if err != nil {
			s.Logger.Error(err)
		}
	})
}

func (s *APIServer) register() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("REGISTER USER REQUEST")
		var temp models.User
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.Logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.Unmarshal(reqBody, &temp)
		if err != nil {
			s.Logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		s.Users.RegisterUser(temp)
	})
}

func (s *APIServer) auth() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("AUTH USER REQUEST")
		var temp models.User
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.Logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.Unmarshal(reqBody, &temp)
		if err != nil {
			s.Logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		s.Users.AuthUser(temp)
	})
}

func (s *APIServer) edit() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Info("EDIT USER REQUEST")
		var temp models.User
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.Logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.Unmarshal(reqBody, &temp)
		if err != nil {
			s.Logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		s.Users.EditUserProps(temp)
	})
}
