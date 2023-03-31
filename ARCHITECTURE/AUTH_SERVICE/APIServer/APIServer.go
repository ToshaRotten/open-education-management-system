package APIServer

import (
	"encoding/json"
	"fmt"
	"github.com/ToshaRotten/open-educaton-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/APIServer/StatisticManager"
	"github.com/ToshaRotten/open-educaton-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/APIServer/UserManager"
	"github.com/ToshaRotten/open-educaton-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/APIServer/UserManager/models"
	"github.com/ToshaRotten/open-educaton-management-system/ARCHITECTURE/AUTH_SERVICE/APIServer/APIServer/config"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	Config    *config.Config
	Logger    *logrus.Logger
	Users     *UserManager.UserManager
	Router    *mux.Router
	Telemetry *StatisticManager.StatisticManager
}

func New() *APIServer {
	var s APIServer
	s.Config = config.New()
	s.Logger = logrus.New()
	s.Router = mux.NewRouter()
	s.Users = UserManager.New()
	s.Telemetry = StatisticManager.New()
	return &s
}

func (s *APIServer) Start(config *config.Config) error {
	s.Config = config
	err := s.configureLogger()
	if err != nil {
		s.Logger.Error(err)
		return err
	}
	s.configureMemory()
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

func (s *APIServer) configureMemory() {

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
	s.Router.HandleFunc("/service/ping", s.ping())
	s.Router.HandleFunc("/user/register", s.register())
	s.Router.HandleFunc("/user/auth", s.auth())
	s.Router.HandleFunc("/user/edit", s.edit())
	s.Router.HandleFunc("/user/list", s.userList())
}

func (s *APIServer) userList() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		//array := make([]models.User, s.Buffer.Count())
		//for i, user := range s.Buffer.Users {
		//	array[i] = user
		//}
		//data, err := json.Marshal(&array)
		//if err != nil {
		//	s.Logger.Error(err)
		//}
		//_, err = w.Write(data)
		//if err != nil {
		//	s.Logger.Error(err)
		//}
	})
}

func (s *APIServer) ping() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ping")
		s.Users.RegisterUser(models.User{
			FirstName: "",
			LastName:  "",
			ThirdName: "",
			Phone:     "",
			Email:     "",
			DOB:       "",
			Role:      "",
			Hash:      "",
		})
	})
}

func (s *APIServer) getTelemetry() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(&s.Telemetry.Last)
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
		//var temp models.User
		//reqBody, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	s.Logger.Error(err)
		//	w.WriteHeader(http.StatusBadRequest)
		//}
		//err = json.Unmarshal(reqBody, &temp)
		//if err != nil {
		//	s.Logger.Error(err)
		//	w.WriteHeader(http.StatusBadRequest)
		//}
		//if s.Buffer.Search(temp) {
		//	s.Logger.Warn("USER IS EXIST")
		//	w.WriteHeader(http.StatusAlreadyReported)
		//} else {
		//	s.Buffer.Add(temp)
		//	s.Logger.Info("USER IS CREATED")
		//	s.Buffer.Print()
		//	w.WriteHeader(http.StatusAccepted)
		//}
	})
}

func (s *APIServer) auth() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//s.Logger.Info("AUTH USER REQUEST")
		//var temp models.User
		//reqBody, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	s.Logger.Error(err)
		//	w.WriteHeader(http.StatusBadRequest)
		//}
		//err = json.Unmarshal(reqBody, &temp)
		//if err != nil {
		//	s.Logger.Error(err)
		//	w.WriteHeader(http.StatusBadRequest)
		//}
		//if s.Buffer.SearchByHash(temp) && s.Buffer.SearchByEmail(temp) {
		//	s.Logger.Trace("IS FOUND")
		//	w.WriteHeader(http.StatusFound)
		//} else {
		//	w.WriteHeader(http.StatusNotFound)
		//}
	})
}

func (s *APIServer) edit() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//s.Logger.Info("EDIT USER REQUEST")
		//var temp models.User
		//reqBody, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	s.Logger.Error(err)
		//	w.WriteHeader(http.StatusBadRequest)
		//}
		//err = json.Unmarshal(reqBody, &temp)
		//if err != nil {
		//	s.Logger.Error(err)
		//	w.WriteHeader(http.StatusBadRequest)
		//}
		//if s.Buffer.SearchByHash(temp) {
		//	//s.Logger.Trace("BEFORE:")
		//	//s.Buffer.Print()
		//	s.Buffer.Replace(temp)
		//	s.Logger.Info("USER IS REPLACED")
		//	//s.Logger.Trace("AFTER:")
		//	//s.Buffer.Print()
		//	w.WriteHeader(http.StatusAccepted)
		//} else {
		//	s.Logger.Warn("USER NOT FOUND")
		//	w.WriteHeader(http.StatusNotFound)
		//}
	})
}
