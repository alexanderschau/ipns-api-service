package main

import (
	"net/http"

	"go.alxs.xyz/ipns-service/server"
)

func main() {
	http.ListenAndServe(":8081", http.HandlerFunc(server.Handler))
}
