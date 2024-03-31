package routes

import (
	"go_api/world/internal/cities/handlers"
	"go_api/world/internal/cities/repositories"
	"go_api/world/internal/cities/routes"

	user_handler "go_api/world/internal/user/handlers"
	user_repositories "go_api/world/internal/user/repositories"
	user_routes "go_api/world/internal/user/routes"
	user_services "go_api/world/internal/user/services"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server up!"))
	})

	r.Mount("/cities", routes.CitiesRouter(initCityHandler()))
	r.Mount("/users", user_routes.UserRouter(initUserHandler()))
}

func initCityHandler() *handlers.CityHandler {
	cr := repositories.NewCityRepository()

	ch := handlers.NewCityHandler(cr)

	return ch
}

func initUserHandler() *user_handler.UserHandler {
	ur := user_repositories.NewUserMongoRepository()
	us := user_services.NewUserService(ur)
	uh := user_handler.NewUserHandler(us)

	return &uh
}
