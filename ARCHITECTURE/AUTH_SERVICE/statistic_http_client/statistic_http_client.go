package statistic_http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/statistic_http_client/logger"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/statistic_http_client/models"
	"net/http"
)

var (
	deployAddr = "185.21.142.92:9908"
	testAddr   = "localhost:9908"
)

type StatisticClient struct {
	enabled bool
	client  http.Client
	logger  *logger.STATLogger
}

func New() *StatisticClient {
	var s StatisticClient
	s.client = http.Client{}
	s.logger = logger.New()
	return &s
}

func (s *StatisticClient) HealthCheck() error {
	req, err := http.NewRequest("POST", "http://"+testAddr+"/health", nil)
	if err != nil {
		s.logger.Log(fmt.Sprintf("Req error: %s", err.Error()))
		s.enabled = false
		return err
	}
	_, err = s.client.Do(req)
	if err != nil {
		s.logger.Log(fmt.Sprintf("Resp error: %s", err.Error()))
		s.enabled = false
		return err
	}
	s.enabled = true
	return nil
}

func (s *StatisticClient) SendLog(message string) {
	if s.enabled {
		var log models.Log
		log.Message = "[AUTH_SERVICE_LOG]: " + message
		data, err := json.Marshal(&log)

		req, err := http.NewRequest("POST", "http://"+testAddr+"/log/send", bytes.NewBuffer(data))
		if err != nil {
			s.logger.Log(fmt.Sprintf("Req error: %s", err.Error()))
		}

		req.Header.Set("Content-Type", "application/json")
		_, err = s.client.Do(req)
		if err != nil {
			s.logger.Log(fmt.Sprintf("Resp error: %s", err.Error()))
		}
	}
}
