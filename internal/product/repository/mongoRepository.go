package repository

import (
	"context"
	"log"

	"github.com/nicewook/mg/internal/product/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepo struct {
	client *mongo.Client
}

var _ ProductRepository = (*MongoProductRepo)(nil)

func NewMongoProductRepo(client *mongo.Client) ProductRepository {
	return &MongoProductRepo{
		client: client,
	}
}

func (m *MongoProductRepo) InsertOne(r entity.ProductInsertOneReq) (entResp entity.ProductInsertOneResp, err error) {
	log.Println("insert one")
	collection := m.client.Database(r.Database).Collection(r.Collection)

	var result *mongo.InsertOneResult
	result, err = collection.InsertOne(context.TODO(), r.Product)
	if err != nil {
		return entResp, err
	}

	entResp.InsertedID = result.InsertedID.(primitive.ObjectID).Hex()
	return entResp, err
}

func (m *MongoProductRepo) FindOne(r entity.ProductFindOneReq) (entResp entity.ProductFindOneResp, err error) {
	collection := m.client.Database(r.Database).Collection(r.Collection)

	filter := bson.D{{Key: "type", Value: r.Type}}
	if err = collection.FindOne(context.TODO(), filter).Decode(&entResp.Product); err != nil {
		log.Println(err)
	}
	log.Println("entResp:", entResp)
	return entResp, err
}

func (m *MongoProductRepo) FindMany(r entity.ProductFindManyReq) (entResp entity.ProductFindManyResp, err error) {
	collection := m.client.Database(r.Database).Collection(r.Collection)

	filter := bson.D{{Key: "type", Value: r.Type}}
	ctx := context.TODO()
	curser, err := collection.Find(ctx, filter)
	if err != nil {
		return entResp, err
	}
	defer curser.Close(ctx)

	for curser.Next(ctx) {
		var product entity.Product
		if err := curser.Decode(&product); err != nil {
			return entResp, err
		}
		entResp.Products = append(entResp.Products, product)
	}

	log.Println("entResp.Products:", entResp.Products)
	return entResp, err
}
