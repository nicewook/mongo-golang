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

func (m *MongoProductRepo) Insert(r entity.ProductInsertReq) (entResp entity.ProductInsertResp, err error) {
	log.Println("insert one")
	collection := m.client.Database(r.Database).Collection(r.Collection)

	var (
		result *mongo.InsertManyResult
		docs   []interface{}
	)
	for _, d := range r.Products {
		docs = append(docs, d)
	}

	result, err = collection.InsertMany(context.TODO(), docs)
	if err != nil {
		return entResp, err
	}

	for _, id := range result.InsertedIDs {
		entResp.InsertedIDs = append(entResp.InsertedIDs, id.(primitive.ObjectID).Hex())
	}
	log.Println("inserted ids:", entResp.InsertedIDs)
	return entResp, err
}

func (m *MongoProductRepo) FindOne(r entity.ProductFindOneReq) (entResp entity.ProductFindOneResp, err error) {
	collection := m.client.Database(r.Database).Collection(r.Collection)

	// limitation: cannot query on nested field
	// in this case we might need for POST with JSON of filter
	filter := bson.D{}
	for k, v := range r.QueryParams {
		filter = append(filter, bson.E{Key: k, Value: v[0]})
	}

	if err = collection.FindOne(context.TODO(), filter).Decode(&entResp.Product); err != nil {
		log.Println(err)
	}
	log.Println("found one document")
	return entResp, err
}

func (m *MongoProductRepo) FindMany(r entity.ProductFindManyReq) (entResp entity.ProductFindManyResp, err error) {
	collection := m.client.Database(r.Database).Collection(r.Collection)

	// limitation: cannot query on nested field
	// in this case we might need for POST with JSON of filter
	filter := bson.D{}
	for k, v := range r.QueryParams {
		filter = append(filter, bson.E{Key: k, Value: v[0]})
	}

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

	log.Println("found document count:", len(entResp.Products))
	return entResp, err
}
