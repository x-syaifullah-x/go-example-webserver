package routes

import (
	"net/http"

	"github.com/x-syaifullah-x/web-server/controller"
)

func MapRoutes(server *http.ServeMux) {
	server.HandleFunc("/", controller.HelloWord())
	server.HandleFunc("/users", controller.Users())
	server.HandleFunc("/index", controller.Indexs())
}
