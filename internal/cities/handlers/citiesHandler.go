package handlers

import (
	"encoding/json"
	"go_api/world/internal/cities/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CityHandler struct {
	repo domain.CityDb
}

func NewCityHandler(repo domain.CityDb) *CityHandler {
	return &CityHandler{repo: repo}
}

func (handler *CityHandler) GetCity(w http.ResponseWriter, r *http.Request) {

	cityID := chi.URLParam(r, "id")

	city, _ := handler.repo.GetCityById(cityID)

	json.NewEncoder(w).Encode(city)
}

func (handler *CityHandler) GetCities(w http.ResponseWriter, r *http.Request) {

	lt := r.URL.Query().Get("lt")
	gt := r.URL.Query().Get("gt")
	countryCode := r.URL.Query().Get("cc")
	district := r.URL.Query().Get("district")

	cities, _ := handler.repo.GetCities(gt, lt, countryCode, district)

	json.NewEncoder(w).Encode(cities)

}
