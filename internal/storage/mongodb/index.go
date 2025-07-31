package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func createIndex(ctx context.Context, coll *mongo.Collection) error {
	indexModel := mongo.IndexModel{
		Keys: bson.D{{"name", 1}},
	}

	_, err := coll.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("index already exists: %w", err)
		}
		return fmt.Errorf("failed to create index: %w", err)
	}

	return nil
}
