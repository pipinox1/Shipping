package stream

import (
	"apitest/internal/domain/model"
	"apitest/internal/tools/publisher"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type streamMongo struct {
	*mongo.Collection
	publisher.Publisher
}

func NewMongoStreamHandler(mongoCollection *mongo.Collection, publisher publisher.Publisher) Stream{
	return &streamMongo{mongoCollection,publisher}
}

func (s streamMongo) StreamDatabaseId()  {
	shipmentStream, err := s.Collection.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		panic(err)
	}
	defer shipmentStream.Close(context.Background())
	changeDoc := struct {
		FullDocument model.Shipment `bson:"fullDocument"`
	}{}

	for shipmentStream.Next(context.TODO()) {
		if err := shipmentStream.Decode(&changeDoc); err != nil {
			panic(err)
		}
		s.Publisher.PublishMessage(changeDoc.FullDocument.Id,"shipment")
	}
}