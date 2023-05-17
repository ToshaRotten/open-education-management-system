package statistic_http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/AUTH_SERVICE/statistic_http_client/models"
	"net/http"
)

var (
	deployAddr = "185.21.142.92:9908"
	testAddr   = "localhost:9908"
)

type StatisticClient struct {
	client http.Client
}

func New() *StatisticClient {
	var s StatisticClient
	s.client = http.Client{}
	return &s
}

func (s *StatisticClient) SendLog(message string) {
	var log models.Log
	log.Message = "[AUTH_SERVICE_LOG]: " + message
	data, err := json.Marshal(&log)

	req, err := http.NewRequest("POST", "http://"+testAddr+"/log/send", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Req error: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	_, err = s.client.Do(req)
	if err != nil {
		fmt.Println("resp error: ", err)
	}
}
