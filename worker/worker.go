package worker

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SpinWorker(port int, id string) error {
	addr := fmt.Sprintf(":%d", port)
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] Landing request", id)
		w.Write([]byte(fmt.Sprintf("Hello from %s", id)))
	})
	err := http.ListenAndServe(addr, r)
	return err
}
