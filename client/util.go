package client

import (
	"net"
)

const (
	ACTIVE   = 1
	INACTIVE = 0
)

func GetControllerStatus() int {
	_, err := net.Listen("tcp", ADDR)
	if err != nil {
		return ACTIVE
	}
	return INACTIVE
}
