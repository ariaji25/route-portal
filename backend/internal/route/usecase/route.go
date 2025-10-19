package usecase

import (
	"context"
	"errors"
	"test/portal/domain"
)

type routeUsecase struct {
	repo domain.RouteItemRepository
}

// Create implements domain.RouteItemUsecase.
func (u *routeUsecase) Create(ctx context.Context, route domain.RouteItem) (*domain.RouteItem, error) {
	existRoute, err := u.repo.GetOne(ctx, route.Name)
	if err != nil {
		if err.Error() != domain.ErrNotFound {
			return nil, err
		}
	}
	if existRoute != nil {
		return nil, errors.New(domain.ErrBadRequest)
	}

	createdRoute, err := u.repo.Create(ctx, route)
	if err != nil {
		return nil, errors.New(domain.ErrInternalServer)
	}

	return createdRoute, nil
}

// Delete implements domain.RouteItemUsecase.
func (u *routeUsecase) Delete(ctx context.Context, name string) error {
	_, err := u.repo.GetOne(ctx, name)
	if err != nil {
		return err
	}
	err = u.repo.Delete(ctx, name)
	if err != nil {
		return errors.New(domain.ErrInternalServer)
	}
	return nil
}

// GetAll implements domain.RouteItemUsecase.
func (u *routeUsecase) GetAll(ctx context.Context) ([]domain.RouteItem, error) {
	routes, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, errors.New(domain.ErrInternalServer)
	}

	return routes, nil
}

// GetOne implements domain.RouteItemUsecase.
func (u *routeUsecase) GetOne(ctx context.Context, name string) (*domain.RouteItem, error) {
	route, err := u.repo.GetOne(ctx, name)
	if err != nil {
		return nil, errors.New(domain.ErrInternalServer)
	}

	return route, nil
}

// Update implements domain.RouteItemUsecase.
func (u *routeUsecase) Update(ctx context.Context, route domain.RouteItem) (*domain.RouteItem, error) {
	_, err := u.repo.GetOne(ctx, route.Name)
	if err != nil {
		return nil, err
	}

	updatedRoute, err := u.repo.Update(ctx, route)
	if err != nil {
		return nil, err
	}

	return updatedRoute, nil

}

func NewRouteUsecase(repo domain.RouteItemRepository) domain.RouteItemUsecase {
	return &routeUsecase{
		repo: repo,
	}
}
