package repository_test

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDB(dsn string) (*mongo.Client, error) {
	return mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
}
func TestProductInsertOne(t *testing.T) {
	// database connect
	client, err := connectDB("mongodb://127.0.0.1:27017")
	if err != nil {
		t.Error("failed to connect")
	}
	t.Log("database server connected")
	_ = client

}
