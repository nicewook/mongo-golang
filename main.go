package main

import (
	"context"
	"log"
	"os"

	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handler "github.com/nicewook/mg/internal/product/handler/http"
	"github.com/nicewook/mg/internal/product/repository"
	"github.com/nicewook/mg/internal/product/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/nicewook/mg/docs"
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

// @title          Golang MongoDB CRUD - Clean Architecture example
// @version        1.0
// @description    This show two things. Golang-MongoDB CRUD server and Clean Architecture.
// @termsOfService http://swagger.io/terms/

// @contact.name  hyunseok.jeong
// @contact.url   http://www.annotation-ai.com
// @contact.email hyunseok.jeong@annotation-ai.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8888
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
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(echoPrometheus.MetricsMiddleware())
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	handler.NewProductHandler(e, productService)

	log.Fatal(e.Start(":8888"))
}
