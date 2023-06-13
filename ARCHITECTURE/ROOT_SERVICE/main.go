package main

import (
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/ROOT_SERVICE/root_http_server"
)

func Init() {

}

func main() {
	//var controller ServiceHelper.ServiceHelper
	serv := root_http_server.New()
	serv.Start()
}
