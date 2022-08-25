package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/nicewok/mg/internal/product/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepo struct {
	client *mongo.Client
}

var _ ProductRepository = (*MongoProductRepo)(nil)

func (m *MongoProductRepo) InsertOne(r entity.ProductInsertOneReq) entity.ProductInsertOneResp {
	coll := m.client.Database(r.Database).Collection(r.Collection)
	result, err := coll.InsertOne(context.TODO(), r)
	if err != nil {
		log.Println(err)
		return entity.ProductInsertOneResp{
			InsertedID: "",
			Message:    "failed to insert",
		}
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	insertedID := result.InsertedID.(primitive.ObjectID)

	return entity.ProductInsertOneResp{
		InsertedID: fmt.Sprintf("v", insertedID),
		Message:    fmt.Sprintf("_id: %v inserted", insertedID),
	}
}
func (m *MongoProductRepo) FindOne(r entity.ProductFindOneReq) (entResp entity.ProductFindOneResp, err error) {
	collection := m.client.Database(r.Database).Collection(r.Collection)

	var doc bson.M
	filter := bson.D{{Key: "type", Value: r.Type}}
	if err = collection.FindOne(context.Background(), filter).Decode(&doc); err != nil {
		log.Println(err)
	}
	log.Println("doc bson.M:", doc)

	return entResp, err
}

func NewMongoProductRepo(client *mongo.Client) ProductRepository {
	return &MongoProductRepo{
		client: client,
	}
}
