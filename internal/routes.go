package routes

import "net/http"

func Get() {

	http.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("testing get"))
		w.WriteHeader(http.StatusOK)
	})

}
