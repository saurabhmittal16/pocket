package worker

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SpinWorker(port int, id string) error {
	addr := fmt.Sprintf(":%d", port)
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hello from %s", id)))
	})
	err := http.ListenAndServe(addr, r)
	return err
}
