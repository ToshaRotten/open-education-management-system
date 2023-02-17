package server

import (
	. "main/loger"
	"main/server/configurator"
)

type APIServer struct {
	Config *configurator.Config
}

//New ..
func New() *APIServer {
	return &APIServer{
		Config: nil,
	}
}

//Start ..
func (s *APIServer) Start() {
	GlobalLoger.Info("Server is started at port:", s.Config.Port)
	GlobalLoger.Info("Server bind addr:", "http://"+s.Config.Host+":"+s.Config.Port)
	s.Configurate()
}

//Configurate ..
func (s *APIServer) Configurate() {
	GlobalLoger.Trace("Configurate APIServer")
	//s.Config.ParseFile()
}
