package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	routedelivery "test/portal/internal/route/delivery/http"
	routeyamlrepository "test/portal/internal/route/repository/yaml"
	routeusecase "test/portal/internal/route/usecase"
	"test/portal/pkg/validations"

	"github.com/go-playground/validator/v10"
)

type App struct{}

func newApp() App {
	return App{}
}

func (*App) start(ctx context.Context) {
	// Initiate all dependencies
	// Initiate repository
	routeRepo := routeyamlrepository.NewRouteYamlRepository()

	// Initiate usecase
	routeUsecase := routeusecase.NewRouteUsecase(routeRepo)

	// Initiate custom validator dependencies
	customValidator := validator.New()
	if err := customValidator.RegisterValidation("is_valid_name", validations.IsValidName); err != nil {
		log.Println("Failed initiate validator is_valid_name", err)
	}
	if err := customValidator.RegisterValidation("is_valid_path", validations.IsValidPath); err != nil {
		log.Println("Failed initiate validator is_valid_path", err)
	}
	if err := customValidator.RegisterValidation("is_valid_host", validations.IsValidHostName); err != nil {
		log.Println("Failed initiate validator is_valid_host", err)
	}
	if err := customValidator.RegisterValidation("is_valid_backend_url", validations.IsValidBackendUrl); err != nil {
		log.Println("Failed initiate validator is_valid_backend_url", err)
	}

	// Initiate delivery
	routedelivery.NewRouteDelivery(ctx, customValidator, routeUsecase)
	// Health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Uk0tMjAyNS0xMC1BTDdRMuKAjAo="))
	})

	fmt.Println("Start The Web Service on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	// Entry point for the API server
	ctx := context.Background()
	app := newApp()
	app.start(ctx)
}
