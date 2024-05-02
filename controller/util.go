package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
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

func GetValue(addr string, key string) ([]byte, error) {
	workerReq := fmt.Sprintf("%s?key=%s", addr, key)

	resp, err := http.Get(workerReq)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func PostValue(addr, key, value string) ([]byte, error) {
	postBody, _ := json.Marshal(map[string]string{
		"key":   key,
		"value": value,
	})
	body := bytes.NewBuffer(postBody)
	resp, err := http.Post(addr, "application/json", body)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	var respBody []byte
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return respBody, nil
}
