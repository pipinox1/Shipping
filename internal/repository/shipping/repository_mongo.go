package shipping

import (
	"apitest/internal/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRepositoryMongo(collection *mongo.Collection)  Repository {
	return &repositoryMongo{Collection: collection}
}

type repositoryMongo struct {
	Collection *mongo.Collection
}

func (r *repositoryMongo) GetById(ctx context.Context, id string) (*model.Shipment, error) {
	panic("implement me")
}

func (r *repositoryMongo) Save(ctx context.Context, shipment *model.Shipment) error {
	_, err := r.Collection.InsertOne(ctx,shipment)
	if err != nil {
		//do something with error
		return err
	}
	return nil
}





