package controller

import (
	"auth_service/service"
	"auth_service/web"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	svc *service.AuthServiceImpl
}

func NewAuthServiceImpl(svc *service.AuthServiceImpl) *AuthControllerImpl {
	return &AuthControllerImpl{
		svc: svc,
	}
}

func (ctrl *AuthControllerImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	reqBody := web.Request{}
	json.NewDecoder(r.Body).Decode(&reqBody)

	token, err := ctrl.svc.GenerateJwt(r.Context(), &reqBody)
	if err != nil {
		w.WriteHeader(400)
		log.Print(err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    token.Refresh,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(2 * time.Hour),
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(token)
}
