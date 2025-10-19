package yaml

import (
	"context"
	"errors"
	"log"
	"test/portal/domain"
	"test/portal/pkg/sliceutils"
	"test/portal/pkg/yamlutils"
)

type routeYamlRepository struct {
	yamlPath string
}

// Create implements domain.RouteItemRepository.
func (r *routeYamlRepository) Create(ctx context.Context, route domain.RouteItem) (*domain.RouteItem, error) {
	existRoutes := make([]domain.RouteItem, 0)
	err := yamlutils.LoadYamlData(r.yamlPath, &existRoutes)
	if err != nil {
		return nil, err
	}
	existRoutes = append(existRoutes, route)
	err = yamlutils.SaveYamlData(r.yamlPath, existRoutes)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &route, nil
}

// Delete implements domain.RouteItemRepository.
func (r *routeYamlRepository) Delete(ctx context.Context, name string) error {
	existRoutes := make([]domain.RouteItem, 0)
	err := yamlutils.LoadYamlData(r.yamlPath, &existRoutes)
	if err != nil {
		return err
	}

	existRoutes = sliceutils.Filter(existRoutes, func(ri domain.RouteItem) bool { return ri.Name != name })

	err = yamlutils.SaveYamlData(r.yamlPath, existRoutes)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements domain.RouteItemRepository.
func (r *routeYamlRepository) GetAll(ctx context.Context) ([]domain.RouteItem, error) {
	existRoutes := make([]domain.RouteItem, 0)
	err := yamlutils.LoadYamlData(r.yamlPath, &existRoutes)
	if err != nil {
		return nil, err
	}
	return existRoutes, nil
}

// GetOne implements domain.RouteItemRepository.
func (r *routeYamlRepository) GetOne(ctx context.Context, name string) (*domain.RouteItem, error) {
	existRoutes := make([]domain.RouteItem, 0)
	err := yamlutils.LoadYamlData(r.yamlPath, &existRoutes)
	if err != nil {
		return nil, err
	}
	existRoutes = sliceutils.Filter(existRoutes, func(ri domain.RouteItem) bool { return ri.Name == name })
	if len(existRoutes) == 0 {
		return nil, errors.New(domain.ErrNotFound)
	}

	route := existRoutes[0]
	return &route, nil
}

// Update implements domain.RouteItemRepository.
func (r *routeYamlRepository) Update(ctx context.Context, route domain.RouteItem) (*domain.RouteItem, error) {
	existRoutes := make([]domain.RouteItem, 0)
	err := yamlutils.LoadYamlData(r.yamlPath, &existRoutes)
	if err != nil {
		return nil, err
	}
	existRoutes = sliceutils.Filter(existRoutes, func(ri domain.RouteItem) bool { return ri.Name != route.Name })
	existRoutes = append(existRoutes, route)

	err = yamlutils.SaveYamlData(r.yamlPath, existRoutes)
	if err != nil {
		return nil, err
	}
	return &route, nil
}

func NewRouteYamlRepository() domain.RouteItemRepository {
	return &routeYamlRepository{
		yamlPath: "./.data/routes.yaml",
	}
}
