package database

type Helper struct {
	Host string
	Port string
}

func New() *Helper {
	return &Helper{
		Host: "",
		Port: "",
	}
}
