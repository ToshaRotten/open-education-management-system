package configurator_http_server

import (
	"fmt"
	"net"
	"strconv"
)

func SearchAvailablePort() string {
	res := ""
	for port := 8000; port < 9000; port++ {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			ln.Close()
			fmt.Printf("Port %d is available\n", port)
			res = strconv.Itoa(port)
			break
		}
	}
	return res
}

func SearchIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			fmt.Println("Server IP address is:", ipnet.IP.String())
		}
	}
}
