package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Register(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
