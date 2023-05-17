package schedule_http_server

import "testing"

func TestNew(t *testing.T) {}

func TestServer_Start(t *testing.T) {
	s := New()
	s.Start()
}
