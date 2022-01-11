package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/ehsun7b/securego/server/rest"
	"github.com/ehsun7b/securego/server/static"
)

func Start() {
	http.HandleFunc("/", handler)
	startServer()
}

func handler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if !strings.HasPrefix(path, "/api") {
		static.Serve(w, req)
	} else {
		rest.Serve(w, req)
	}

}

func startServer() {
	log.Println("server started")
	error := http.ListenAndServe(":9090", nil)

	if error != nil {
		log.Println(error)
	}
}
