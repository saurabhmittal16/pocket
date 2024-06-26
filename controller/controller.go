package controller

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"sync"

	"github.com/saurabhmittal16/pocket/worker"
)

// total slots for hash space
const TOTAL_SLOTS = 1e10

var lock = &sync.Mutex{}

type workerNode struct {
	Id   string
	Port int
	Hash uint64
}

type controllerNode struct {
	ring []workerNode
}

var instance *controllerNode

func GetControllerInstance() *controllerNode {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		// workers[i] is at space[i] in the hash space
		instance = &controllerNode{ring: []workerNode{}}
	}
	return instance
}

func (c *controllerNode) CreateWorkers(numWorkers int) error {
	count := len(c.ring)
	ports, err := GetAvailablePorts(numWorkers)

	if err != nil {
		return err
	}

	for i := 0; i < numWorkers; i++ {
		port := ports[i]

		// create workerNode instance
		id := fmt.Sprintf("Worker#%05d", count+i+1)
		w := workerNode{Id: id, Port: port, Hash: hash(id)}

		// spin up REST server for worker node
		go worker.SpinWorker(port, id)

		// save the workerNode instance
		c.addToRing(w)

		// TODO: Get error from SpinWorker and log accordingly
		log.Printf("Successfuly started worker node (Id = %s) at %d", id, port)
	}
	return nil
}

func (c *controllerNode) addToRing(w workerNode) {
	c.ring = append(c.ring, w)
	sort.SliceStable(c.ring, func(i, j int) bool {
		return c.ring[i].Hash < c.ring[j].Hash
	})
}

func (c *controllerNode) FindWorker(key string) (workerNode, error) {
	if len(c.ring) == 0 {
		return workerNode{}, fmt.Errorf("no worker nodes running")
	}

	// get hash value for cache key
	h := hash(key)

	// binary search the space
	i := sort.Search(len(c.ring), func(i int) bool {
		return c.ring[i].Hash > h
	})

	// Check if the key was actually found (might be at the end)
	if i < len(c.ring) && c.ring[i].Hash > h {
		return c.ring[i], nil // Index of struct with score greater than target
	}
	// input hash is greater than all elements, return first index
	return c.ring[0], nil
}

func hash(key string) uint64 {
	// generate the hash of key
	sum := sha256.Sum256([]byte(key))

	// use the first 32 bytes to generate integer position
	data := binary.BigEndian.Uint64(sum[:32])

	// take mod with TOTAL SLOTS space
	return data % TOTAL_SLOTS
}
