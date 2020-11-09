package main

import (
	orders2 "apitest/internal/domain/orders"
	"apitest/internal/domain/shipping"
	user2 "apitest/internal/domain/user"
	"apitest/internal/repository/orders"
	shippingrepository "apitest/internal/repository/shipping"
	"apitest/internal/repository/user"
	"apitest/internal/rest"
	"apitest/internal/stream"
	publisher2 "apitest/internal/tools/publisher"
	"apitest/internal/tools/restclient"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	database := os.Getenv("PERSISTENCE_ARCHITECTURE")
	//Rabbit Connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic("error connecting rabbit")
	}

	ch, err := conn.Channel()
	if err != nil {
		panic("error open rabbit channel")
	}
	publisher := publisher2.NewPublisher(ch)

	var streamDb stream.Stream
	var shipmentRepository shipping.Repository
	if database == "DynamoDB"{
		//Create DynamoDb configuration
		sess := session.Must(session.NewSession(&aws.Config{
			Region:   aws.String("us-east-2"),
			Credentials: credentials.NewSharedCredentials("C:\\Users\\Eduardo\\.aws\\credentials", "pipinox2706"),
		}))
		db := dynamodb.New(sess)
		streamDb = stream.NewDynamoStreamHandler(sess,db,publisher)
		shipmentRepository =  shippingrepository.NewRepositoryDynamo(db, "shipment")
	}else if database == "Mongo"{
		clientOpts := options.Client().ApplyURI("mongodb+srv://shipment:shipment@cluster0.rbn9k.mongodb.net/shipment?retryWrites=true&w=majority")
		client, err := mongo.Connect(context.TODO(), clientOpts)
		if err != nil {
			log.Fatal(err)
		}
		collection := client.Database("shipment").Collection("shipment")
		shipmentRepository = shippingrepository.NewRepositoryMongo(collection)
		streamDb = stream.NewMongoStreamHandler(collection,publisher)
	}else {
		panic("not database selected")
	}

	//Repository
	userRepo := user.NewRepository(restclient.NewRestClient(2 * time.Second))
	orderRepo := orders.NewRepository(restclient.NewRestClient(5 * time.Second))

	//Service
	orderService := orders2.NewService(orderRepo)
	userService := user2.NewService(userRepo)
	shippingService := shipping.NewService(shipmentRepository, orderService, userService)

	//Handlers
	sh := rest.NewShippingHandler(shippingService)
	handlerDb := stream.NewStreamHandler(streamDb)

	//Subscriber Handler
	go func() {
		handlerDb.HandlerDatabaseMessage()
	}()

	//Start App
	router := chi.NewRouter()
	rest.RegisterShippingRoute(router, sh)

	log.Print("Availability routes")
	for _, a := range router.Routes() {
		for _, b := range a.SubRoutes.Routes() {
			log.Print(fmt.Sprint(strings.ReplaceAll(a.Pattern, "/*", ""), b.Pattern))
		}

	}
	port := ":8020"
	log.Print(fmt.Sprint("Starting server at port", port))
	log.Fatal(http.ListenAndServe(port, router))

}
