package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ehsun7b/securego/server/service/status"
)

func serviceStatus(w http.ResponseWriter, req *http.Request) {
	log.Println("status service ...")

	sta := status.Status()
	bytes, err := json.Marshal(sta)
	if err != nil {
		fmt.Println(err)
		resp500(w, err)
		return
	}

	respOkJson(w, bytes)

	updates, ok := req.URL.Query()["update"]
	if ok && updates[0] == "true" {
		status.Update()
	}
}
