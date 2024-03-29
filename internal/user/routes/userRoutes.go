package routes

import (
	"go_api/world/internal/user/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(uh *handlers.UserHandler) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", uh.RegisterUser)

	return r

}
