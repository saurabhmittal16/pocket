package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type request struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func SpinBalancer(port int) error {
	addr := fmt.Sprintf(":%d", port)
	r := chi.NewRouter()

	// landing route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[REST][LB] Landing")
		w.Write([]byte("Hello from Controller"))
	})

	r.Route("/cache", func(r chi.Router) {
		r.Get("/", getValue)
		r.Post("/", postValue)
	})

	err := http.ListenAndServe(addr, r)
	return err
}

func getValue(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	log.Printf("[REST][LB] GET: %v", key)

	// if no key, can't redirect, throw error
	if len(key) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No key received!"))
		return
	}

	workerNode, err := getWorkerNode(key)
	if err != nil {
		writeError(err, w)
	}

	workerAddr := fmt.Sprintf("http://localhost:%d", workerNode.Port)
	log.Printf("[REST][LB] Send request to %s (%s)", workerNode.Id, workerAddr)

	// get response from worker
	body, err := GetValue(workerAddr, key)
	if err != nil {
		writeError(err, w)
	}

	// respond
	w.Write(body)
}

func postValue(w http.ResponseWriter, r *http.Request) {
	var requestObj request
	json.NewDecoder(r.Body).Decode(&requestObj)
	key := requestObj.Key
	val := requestObj.Value
	log.Printf("[REST][LB] POST: %v, %v", key, val)

	workerNode, err := getWorkerNode(key)
	if err != nil {
		writeError(err, w)
	}

	// redirect request to worker addr
	log.Printf("[REST][LB] Send request to %s (%d)", workerNode.Id, workerNode.Port)
	workerAddr := fmt.Sprintf("http://localhost:%d", workerNode.Port)

	// get reponse from worker
	body, err := PostValue(workerAddr, key, val)
	if err != nil {
		writeError(err, w)
	}

	// respond
	w.Write(body)
}

func getWorkerNode(key string) (workerNode, error) {
	// get controller for redirecting
	c := GetControllerInstance()

	// get worker node as per key
	workerNode, err := c.FindWorker(key)

	return workerNode, err
}

func writeError(err error, w http.ResponseWriter) {
	log.Printf("[REST][LB] Error: %v", err.Error())

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Something went wrong!"))
}
