package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/nicewok/mg/internal/product/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepo struct {
	client *mongo.Client
}

var _ ProductRepository = (*MongoProductRepo)(nil)

func (m *MongoProductRepo) InsertOne(r entity.ProductInsertReq) entity.ProductInsertResp {
	coll := m.client.Database(r.Database).Collection(r.Collection)
	result, err := coll.InsertOne(context.TODO(), r)
	if err != nil {
		log.Println(err)
		return entity.ProductInsertResp{
			InsertedID: "",
			Message:    "failed to insert",
		}
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	insertedID := result.InsertedID.(primitive.ObjectID)

	return entity.ProductInsertResp{
		InsertedID: fmt.Sprintf("v", insertedID),
		Message:    fmt.Sprintf("_id: %v inserted", insertedID),
	}
}

func NewMongoProductRepo(client *mongo.Client) ProductRepository {
	return &MongoProductRepo{
		client: client,
	}
}
