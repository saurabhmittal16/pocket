package controller

import (
	"fmt"
	"net"
)

const START_RANGE int = 8000
const END_RANGE int = 9000

func GetAvailablePort() (int, error) {
	for port := START_RANGE; port <= END_RANGE; port++ {
		address := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", address)
		if err == nil {
			listener.Close()
			return port, nil
		}
	}
	return 0, fmt.Errorf("no available port")
}
