package shipping

import (
	"apitest/internal/domain/model"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func NewRepositoryDynamo(db *dynamodb.DynamoDB,tableName string)  Repository {
	return &repositoryDynamo{DynamoDB: db,TableName: tableName}
}

type repositoryDynamo struct {
	DynamoDB *dynamodb.DynamoDB
	TableName string
}

func (r *repositoryDynamo) GetById(ctx context.Context, id string) (*model.Shipment, error) {
	panic("implement me")
}

func (r *repositoryDynamo) Save(ctx context.Context, shipment *model.Shipment) error {
	movieAVMap, err := dynamodbattribute.MarshalMap(shipment)
	if err != nil {
		return err
	}

	paramsPut := &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item:      movieAVMap,
	}

	_, err = r.DynamoDB.PutItem(paramsPut)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return err
	}

	return nil
}




