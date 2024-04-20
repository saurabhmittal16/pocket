package worker

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SpinWorker(port int) error {
	addr := fmt.Sprintf(":%d", port)
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	err := http.ListenAndServe(addr, r)
	return err
}
