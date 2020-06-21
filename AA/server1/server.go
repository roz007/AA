package server1

import (
	"AA/server1/handler"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes() *chi.Mux {

	r := chi.NewRouter()
	r.Post("/server1", handler.HanderSlashServer1Route)

	return r
}

func StartServer1() {
	serverAddr := ":8001"
	r := SetupRoutes()

	http.ListenAndServe(serverAddr, r)

}
