package main

import (
	"fmt"
	"runtime"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET /")
		reply := fmt.Sprintf("{\"runtime\": \"%v\",\"status\":\"ok\"}", runtime.Version())
		w.Write([]byte(reply))
	})

	fmt.Println("server is running (port=1300)")
	http.ListenAndServe(":1300", r)
}
