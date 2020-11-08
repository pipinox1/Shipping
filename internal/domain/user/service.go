package user

import (
	"apitest/internal/domain/model"
	"context"
	"time"
)


type Service struct {
}

func NewService() *Service {
	return &Service{
	}
}

func (s *Service) ValidateUser(ctx context.Context,shipment *model.Shipment)error{
	//Aca podriamos validar contra la api de usuarios si los user tanto vendedor como comprador existen y si son del mismo pais por ej
	time.Sleep(5*time.Second)
	return nil
}