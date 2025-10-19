package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"test/portal/domain"
	"test/portal/pkg/httputils"

	"github.com/go-playground/validator/v10"
)

type RouteDelivery struct {
	usecase  domain.RouteItemUsecase
	validate *validator.Validate
}

func NewRouteDelivery(
	ctx context.Context,
	validate *validator.Validate,
	usecase domain.RouteItemUsecase) *RouteDelivery {

	handler := &RouteDelivery{
		usecase:  usecase,
		validate: validate,
	}

	http.HandleFunc("/routes/", func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Check for path
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")
		log.Println("Path parts", parts)
		// If the path contain path param for route by name handle it to specific func handler

		if len(parts) == 1 && parts[0] == "routes" {
			// /routes
			switch r.Method {
			case http.MethodGet:
				handler.GetAll(ctx, w, r)
			case http.MethodPost:
				handler.Create(ctx, w, r)
			default:
				w.WriteHeader(http.StatusOK)
			}
			return
		}

		if len(parts) == 2 && parts[0] == "routes" {
			// /routes/{name}
			switch r.Method {
			case http.MethodGet:
				handler.GetOne(ctx, w, r)
			case http.MethodPut:
				handler.Update(ctx, w, r)
			case http.MethodDelete:
				handler.Delete(ctx, w, r)
			default:
				w.WriteHeader(http.StatusOK)
			}
			return
		}

		// Anything else is 404
		http.NotFound(w, r)
	})

	return handler
}

func NewTestRouteDelivery(
	ctx context.Context,
	validate *validator.Validate,
	usecase domain.RouteItemUsecase) *RouteDelivery {

	handler := &RouteDelivery{
		usecase:  usecase,
		validate: validate,
	}

	return handler
}

func (h *RouteDelivery) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	route := &domain.RouteItem{}
	err := httputils.ValidateAndUnmarshal(r, h.validate, route)
	if err != nil {
		httputils.WriteErrorResponse(w, err)
		return
	}

	createdRoute, err := h.usecase.Create(ctx, *route)
	if err != nil {
		httputils.WriteErrorResponse(w, err)
		return
	}

	httputils.WriteSuccessResponse(w, createdRoute)
}

func (h *RouteDelivery) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// Expect the path param on the third index position on the URL
	routeName := httputils.GetPathParamByPathPosition(r, 2)
	if routeName == nil {
		httputils.WriteErrorResponse(w, errors.New(domain.ErrBadRequest))
		return
	}

	route := &domain.RouteItem{}
	err := httputils.ValidateAndUnmarshal(r, h.validate, route)
	if err != nil {
		httputils.WriteErrorResponse(w, err)
		return
	}

	if *routeName != route.Name {
		httputils.WriteErrorResponse(w, errors.New(domain.ErrBadRequest+" :Unable to change the name for the route"))
		return
	}

	createdRoute, err := h.usecase.Update(ctx, *route)
	if err != nil {
		httputils.WriteErrorResponse(w, err)
		return
	}

	httputils.WriteSuccessResponse(w, createdRoute)
}

func (h *RouteDelivery) GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	routes, err := h.usecase.GetAll(ctx)
	if err != nil {
		httputils.WriteErrorResponse(w, errors.New(domain.ErrInternalServer))
		return
	}

	httputils.WriteSuccessResponse(w, routes)
}

func (h *RouteDelivery) GetOne(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// Expect the path param on the third index position on the URL
	routeName := httputils.GetPathParamByPathPosition(r, 2)
	if routeName == nil {
		httputils.WriteErrorResponse(w, errors.New(domain.ErrBadRequest))
		return
	}

	route, err := h.usecase.GetOne(ctx, *routeName)
	if err != nil {
		httputils.WriteErrorResponse(w, err)
	}

	httputils.WriteSuccessResponse(w, route)
}

func (h *RouteDelivery) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// Expect the path param on the third index position on the URL
	routeName := httputils.GetPathParamByPathPosition(r, 2)
	if routeName == nil {
		httputils.WriteErrorResponse(w, errors.New(domain.ErrBadRequest))
		return
	}

	err := h.usecase.Delete(ctx, *routeName)
	if err != nil {
		httputils.WriteErrorResponse(w, err)
		return
	}

	httputils.WriteSuccessResponse(w, nil)
}
