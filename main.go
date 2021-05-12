package main

import (
	"fmt"
	"net/http"

	"go.alxs.xyz/ipns-service/api"
)

func main() {
	http.ListenAndServe(":8081", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	cid := r.URL.Query().Get("cid")
	switch r.URL.Path {
	case "/create":
		if cid == "" {
			sendErr(w, "")
			return
		}
		key, err := api.GenKey(cid)
		if err != nil {
			sendErr(w, "")
			return
		}
		sendOK(w, fmt.Sprintf("%s,%s", key.Name, key.Id))
		return
	case "/rm":
		if token == "" || api.RmKey(token) != nil {
			sendErr(w, "")
			return
		}
		sendOK(w, "")
		return
	case "/update":
		if token == "" || cid == "" || api.UpdateKey(token, cid) != nil {
			sendErr(w, "")
			return
		}
		sendOK(w, "")
		return
	}
	sendErr(w, "")
}

func sendOK(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
}

func sendErr(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusForbidden)
}
