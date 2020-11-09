package shipping

import (
	"apitest/internal/domain/model"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func NewRepositoryDynamo(db *dynamodb.DynamoDB, tableName string) Repository {
	return &repositoryDynamo{DynamoDB: db, TableName: tableName}
}

type repositoryDynamo struct {
	DynamoDB  *dynamodb.DynamoDB
	TableName string
}

func (r *repositoryDynamo) GetById(ctx context.Context, id string) (*model.Shipment, error) {
	result, err := r.DynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(r.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if result.Item == nil {
		//do something
	}
	shipment := &model.Shipment{}

	err = dynamodbattribute.UnmarshalMap(result.Item, shipment)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return shipment,nil
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
