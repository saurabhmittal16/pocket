package controller

import (
	"fmt"
	"net"
)

const START_RANGE int = 8000
const END_RANGE int = 9000

func GetAvailablePorts(num int) ([]int, error) {
	ports := make([]int, 0)
	for port := START_RANGE; port <= END_RANGE; port++ {
		address := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", address)
		if err == nil {
			listener.Close()
			ports = append(ports, port)

			if len(ports) == num {
				return ports, nil
			}
		}
	}

	if len(ports) < num {
		return nil, fmt.Errorf("couldn't get %d ports", num)
	}
	return nil, fmt.Errorf("something unexpected happen")
}
