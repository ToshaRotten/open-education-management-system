package statistic_http_client

import "testing"

func TestStatisticClient_SendLog(t *testing.T) {
	cli := New()
	cli.SendLog("logloglog")
}
