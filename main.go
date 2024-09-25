package main

import (
	"github.com/x-syaifullah-x/web-server/routes"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	routes.MapRoutes(server)
	http.ListenAndServe(":8080", server)
}
