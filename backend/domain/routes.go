package domain

import "context"

type RouteItem struct {
	Name    string `json:"name" validate:"required,min=3,max=32,no_whitespace"`
	Host    string `json:"host" validate:"required,min=5,is_valid_host"`
	Path    string `json:"path" validate:"required,is_valid_path"`
	Backend string `json:"backend" validate:"required,min=5,is_valid_backend_url"`
	Enabled *bool  `json:"enabled" validate:"required"`
}

type RouteItemRepository interface {
	Create(ctx context.Context, route RouteItem) (*RouteItem, error)
	Update(ctx context.Context, route RouteItem) (*RouteItem, error)
	GetAll(ctx context.Context) ([]RouteItem, error)
	GetOne(ctx context.Context, name string) (*RouteItem, error)
	Delete(ctx context.Context, name string) error
}

type RouteItemUsecase interface {
	Create(ctx context.Context, route RouteItem) (*RouteItem, error)
	Update(ctx context.Context, route RouteItem) (*RouteItem, error)
	GetAll(ctx context.Context) ([]RouteItem, error)
	GetOne(ctx context.Context, name string) (*RouteItem, error)
	Delete(ctx context.Context, name string) error
}
