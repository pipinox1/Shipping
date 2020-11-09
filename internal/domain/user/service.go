package user

import (
	"apitest/internal/domain/model"
	"apitest/internal/tools/errors"
	"context"
	"time"
)

type Service interface {
	ValidateUser(ctx context.Context, shipment *model.Shipment) error
}

type Repository interface {
	GetById(ctx context.Context, id int64) (*model.User, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repo: repository,
	}
}

func (s *service) ValidateUser(ctx context.Context, shipment *model.Shipment) error {
	user, err := s.repo.GetById(ctx, shipment.Seller.Id)
	if err != nil {
		return err
	}
	if !user.AllowShipment {
		return errors.NewBusinessError("The user does not have permission to send products")
	}
	//Aca podriamos validar contra la api de usuarios si los user tanto vendedor como comprador existen y si son del mismo pais por ej
	time.Sleep(1 * time.Second)
	return nil
}
