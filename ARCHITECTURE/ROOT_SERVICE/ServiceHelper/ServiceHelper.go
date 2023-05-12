package ServiceHelper

import (
	"github.com/Tosha_Rotten/open-education-management-system/ARCHITECTURE/ROOT_SERVICE/ServiceHelper/service"
	"net/http"
)

// ServiceHelper - struct used for control any service or services
type ServiceHelper struct {
	Client http.Client
}

// Restart - restarts an any service
func (sh *ServiceHelper) Restart(sv service.Service) {

}

// Start - starts an any service
func (sh *ServiceHelper) Start(sv service.Service) {

}

// Ping - send PING request to server
func (sh *ServiceHelper) Ping(sv service.Service) {

}

// GetTelemetry - GetTelemetry from service
func (sh *ServiceHelper) GetTelemetry(sv service.Service) {

}
