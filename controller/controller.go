package controller

import (
	"sync"
)

var lock = &sync.Mutex{}

type worker struct {
	id int
}

type controller struct {
	workers []worker
}

var instance *controller

func GetInstance() *controller {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		instance = &controller{workers: []worker{}}
	}
	return instance
}

func (c *controller) CreateWorkers(numWorkers int) error {
	count := len(c.workers)
	for i := 0; i < numWorkers; i++ {
		id := count + i
		w := worker{id: id}
		c.workers = append(c.workers, w)
	}
	return nil
}
