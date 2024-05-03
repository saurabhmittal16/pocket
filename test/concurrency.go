package main

import (
	"log"
	"sync"
	"time"

	"github.com/saurabhmittal16/pocket/controller"
)

var addr = "http://localhost:3001/cache"

func setName(value string) {
	controller.PostValue(addr, "name", value)
}

func getName() {
	resp, _ := controller.GetValue(addr, "name")
	log.Print(string(resp))
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		setName("saurabh")
	}()

	go func() {
		defer wg.Done()
		setName("mittal")
	}()

	time.Sleep(time.Second * 2)

	getName()
}
