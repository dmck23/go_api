package routes

import (
	"go_api/world/internal/cities/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CitiesRouter(ch *handlers.CityHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/", ch.GetCities)
	r.Get("/{id}", ch.GetCity)

	return r
}
