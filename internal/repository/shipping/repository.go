package shipping

import (
	"apitest/internal/domain/model"
	"context"
)

type Repository interface {
	GetById(ctx context.Context, id string) (*model.Shipment, error)
	Save(ctx context.Context, shipment *model.Shipment) error
}
