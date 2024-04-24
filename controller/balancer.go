package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SpinBalancer(port int) error {
	addr := fmt.Sprintf(":%d", port)
	r := chi.NewRouter()

	// landing route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Controller"))
	})

	r.Route("/cache", func(r chi.Router) {
		r.Get("/", getValue)
	})

	err := http.ListenAndServe(addr, r)
	return err
}

func getValue(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	log.Printf("[REST] GET: %v", key)

	c := GetInstance()
	workerNode, err := c.FindWorker(key)

	if err != nil {
		log.Print(err.Error())
	} else {
		log.Printf("[REST] Redirect to %s (:%d)", workerNode.Id, workerNode.Port)
	}

	if len(key) == 0 {
		w.Write([]byte("No key found!"))
	} else {
		w.Write([]byte(fmt.Sprintf("GET %s", key)))
	}
}