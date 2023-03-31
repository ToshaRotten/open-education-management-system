package service

// Service struct
type Service struct {
	Host string
	Port string
}

// New - creates new instance of service
func New() *Service {
	return &Service{
		Host: "",
		Port: "",
	}
}
