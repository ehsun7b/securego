package rest

import (
	"fmt"
	"log"
	"net/http"
)

func Serve(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	log.Println(path)

	if path == "/api/status" {
		serviceStatus(w, req)
		return
	}

	serviceNotFound(w, req)
}

func serviceNotFound(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "service not found 404")
}

func resp500(w http.ResponseWriter, error error) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, error)
}

func respOkJson(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
