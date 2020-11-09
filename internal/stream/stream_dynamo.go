package stream

import (
	"apitest/internal/tools/publisher"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/urakozz/go-dynamodb-stream-subscriber/stream"
)

type streamDynamo struct {
	*session.Session
	*dynamodb.DynamoDB
	publisher.Publisher
}

func (s streamDynamo) StreamDatabaseId()  {
	streamSvc := dynamodbstreams.New(s.Session, &aws.Config{
		Region:   aws.String("us-east-2"),
		Credentials: credentials.NewSharedCredentials("C:\\Users\\Eduardo\\.aws\\credentials", "pipinox2706"),
	})

	streamSubscriber := stream.NewStreamSubscriber(s.DynamoDB, streamSvc, "shipment")
	channelRecord, errCh := streamSubscriber.GetStreamData()

	go func(errCh <-chan error) {
		for err := range errCh {
			fmt.Println("Stream Subscriber error: ", err)
		}
	}(errCh)
	go func() {
		for record := range channelRecord {
			for _, v := range record.Dynamodb.Keys {
				fmt.Println("ID:     ", *v.S)
				s.Publisher.PublishMessage(*v.S,"shipment")
			}
		}
	}()
}

func NewDynamoStreamHandler(sess *session.Session,db *dynamodb.DynamoDB, publisher publisher.Publisher) Stream{
	return &streamDynamo{sess,db,publisher}
}
