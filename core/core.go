package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type key struct {
	title string
}

type value struct {
	value any
}

type Node struct {
	id    uuid.UUID
	elems map[key]value
	port  int
}

func (node Node) Start() (*gin.Engine, error) {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": node.id,
		})
	})
	address := fmt.Sprintf(":%d", node.port)
	router.Run(address)
	log.Print(fmt.Sprintf("Running node %d at port: %d", node.id, node.port))
	return router, nil
}

func (node Node) Get(_key string) any {
	val := node.elems[key{_key}]
	return val.value
}
func (node Node) Put(_key string, _value any) {
	node.elems[key{_key}] = value{value: _value}
}

func CreateNode() *Node {
	portNumber, err := GetAvailablePort()
	if err == nil {
		node := Node{uuid.New(), map[key]value{}, portNumber}
		return &node
	}
	return nil
}
