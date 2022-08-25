package main

import (
	"context"
	"log"
	"os"

	"github.com/labstack/echo"
	handler "github.com/nicewok/mg/internal/product/handler/http"
	"github.com/nicewok/mg/internal/product/repository"
	"github.com/nicewok/mg/internal/product/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoClient() (*mongo.Client, error) {
	uri := "mongodb://127.0.0.1:27017"
	if os.Getenv("DATABASE_URL") != "" {
		uri = os.Getenv("DATABASE_URL")
	}
	return getMongoClientByURI(uri)
}

func getMongoClientByURI(uri string) (*mongo.Client, error) {

	opt := options.Client().ApplyURI(uri).SetMaxPoolSize(5)
	client, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		return client, err
	}
	err = client.Ping(context.Background(), nil)
	return client, err
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("start")

	// MongoDB connection
	client, err := getMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	productRepository := repository.NewMongoProductRepo(client) // get repository interface
	productService := service.NewProductSvc(productRepository)  // get service interface

	e := echo.New()
	handler.NewProductHandler(e, productService)

	log.Fatal(e.Start(":8888"))
}
