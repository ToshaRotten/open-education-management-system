package buffer

type Config struct {
	Id          int    `json:"id" yaml:"id"`
	ServiceName string `json:"service_name" yaml:"service_name"`
	Host        string `json:"host" yaml:"host"`
	Port        string `json:"port" yaml:"port"`
	Scheme      string `json:"scheme" yaml:"scheme"`
	BufferSize  int    `json:"buffer_size" yaml:"buffer_size"`
}
