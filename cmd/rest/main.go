package main

import (
	"fmt"
	"go_api/world/internal/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := run(); err != nil {
		fmt.Println(fmt.Errorf("error - server failed to start. err: %v", err))
	}
}

func run() error {
	r := chi.NewRouter()
	port := viper.Get("server.port")

	routes.Routes(r)

	fmt.Println("Server Starting")

	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
