package orders

import (
	"apitest/internal/domain/model"
	"context"
)

type Repository interface {
	GetById(ctx context.Context,id string) (*model.Order,error)
}

type Service struct {
	repo Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repo: repository,
	}
}

func (s *Service) ValidateOrder(ctx context.Context,shipment *model.Shipment) error{
	_,err:=s.repo.GetById(ctx,shipment.OrderId)
	if err != nil {
		return err
	}
	//Some business logic
	return nil
}
