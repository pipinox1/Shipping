package orders

import (
	"apitest/internal/domain/model"
	"apitest/internal/tools/errors"
	"context"
)

type Repository interface {
	GetById(ctx context.Context, id string) (*model.Order, error)
}

type Service interface {
	ValidateOrder(ctx context.Context, shipment *model.Shipment) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repo: repository,
	}
}

func (s *service) ValidateOrder(ctx context.Context, shipment *model.Shipment) error {
	_, err := s.repo.GetById(ctx, shipment.OrderId)
	//Verify if order exist
	if err != nil {
		if errorCast, ok := err.(*errors.RestClientError); ok {
			if errorCast.StatusCode == 404 {
				return errors.NewBusinessError("order does not exist")
			}
			return err
		}

	}
	//Some business logic
	return nil
}
