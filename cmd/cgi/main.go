package cgi

import (
	"net/http"
	"net/http/cgi"

	"go.alxs.xyz/ipns-service/server"
)

func main() {
	cgi.Serve(http.HandlerFunc(server.Handler))
}
