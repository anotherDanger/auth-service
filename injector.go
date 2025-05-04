//go:build wireinject
// +build wireinject

package main

import (
	"auth_service/controller"
	"auth_service/service"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var ServerSet = wire.NewSet(
	NewRouter,
	NewServer, wire.Bind(new(http.Handler), new(*httprouter.Router)),
	service.NewAuthServiceImpl, wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
	controller.NewAuthServiceImpl, wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
)

func InitServer() *http.Server {
	wire.Build(ServerSet)
	return nil
}
