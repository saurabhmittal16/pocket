package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/saurabhmittal16/pocket/core"
)

type request struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var workerId string
var store *core.Store

func SpinWorker(port int, id string) error {
	workerId = id
	addr := fmt.Sprintf(":%d", port)
	r := chi.NewRouter()

	r.Get("/", getValue)
	r.Post("/", postValue)

	store = core.CreateStore(id)
	err := http.ListenAndServe(addr, r)
	return err
}

func getValue(w http.ResponseWriter, r *http.Request) {
	// key will be non-null since request was redirected as per key
	key := r.URL.Query().Get("key")
	log.Printf("[REST][%s] GET: %v", workerId, key)

	// read value from store
	val := store.Get(key)

	// write value
	w.Write([]byte(val))
}

func postValue(w http.ResponseWriter, r *http.Request) {
	var requestObj request
	json.NewDecoder(r.Body).Decode(&requestObj)
	key := requestObj.Key
	val := requestObj.Value

	log.Printf("[REST][%s] POST: %v, %v", workerId, key, val)

	// put value to store
	store.Put(key, val)

	// respond with 200
	w.WriteHeader(http.StatusOK)
}
