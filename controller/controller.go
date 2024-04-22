package controller

import (
	"log"
	"sync"

	"github.com/saurabhmittal16/pocket/worker"
)

var lock = &sync.Mutex{}

type workerNode struct {
	id   int
	port int
}

type controllerNode struct {
	workers []workerNode
}

var instance *controllerNode

func GetInstance() *controllerNode {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		instance = &controllerNode{workers: []workerNode{}}
	}
	return instance
}

func (c *controllerNode) CreateWorkers(numWorkers int) error {
	count := len(c.workers)
	ports, err := GetAvailablePorts(numWorkers)

	if err != nil {
		return err
	}

	for i := 0; i < numWorkers; i++ {
		port := ports[i]

		// create workerNode instance
		id := count + i
		w := workerNode{id: id, port: port}

		// spin up REST server for worker node
		go worker.SpinWorker(port)

		// save the workerNode instance
		c.workers = append(c.workers, w)

		log.Printf("Successfuly started worker node (Id = %d) at %d", id, port)
	}
	return nil
}
