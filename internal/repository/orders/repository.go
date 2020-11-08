package orders

import (
	"apitest/internal/domain/model"
	"apitest/internal/tools/restclient"
	"context"
	"fmt"
)

var requestUlr =  "localhost:3004/v1/orders/%s"

type Repository interface {
	GetById(ctx context.Context,id string) (*model.Shipment,error)
}

func NewRepositoryDynamo(restClien restclient.RestClient)  Repository {
	return &repository{}
}

type repository struct {
	restclient.RestClient
}

func (r *repository) GetById(ctx context.Context, id string) (*model.Shipment, error) {

	url := fmt.Sprintf(requestUlr,id)

	r.DoGet(ctx,url,)
}

