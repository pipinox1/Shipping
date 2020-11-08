package main

import (
	"apitest/internal/domain/shipping"
	shippingrepository "apitest/internal/repository/shipping"
	"apitest/internal/rest"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"strings"
)

func main() {

	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("sa-east-1"),
		Endpoint: aws.String("http://127.0.0.1:8000"),
		//EndPoint: aws.String("https://dynamodb.us-east-1.amazonaws.com"),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)
	shippingrepository.NewRepositoryDynamo(db,"shipment")


	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", "localhost", 27017))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("shipment").Collection("shipment")
	mongoRepo := shippingrepository.NewRepositoryMongo(collection)


	shippingService := shipping.NewService(mongoRepo)
	sh := rest.NewShippingHandler(shippingService)

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
