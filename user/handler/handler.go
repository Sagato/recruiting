package user_handler

import "github.com/gorilla/mux"

type Handler interface {
	SetupRouter(router *mux.Router) *mux.Router
}
