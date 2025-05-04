package main

import (
	"auth_service/controller"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(ctrl *controller.AuthControllerImpl) *httprouter.Router {
	r := httprouter.New()
	r.POST("/v1/auth", ctrl.Register)
	return r
}

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}
}

func main() {
	server := InitServer()
	err := server.ListenAndServe()
	if err != nil {
		log.Print(err)
	}
}
