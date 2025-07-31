package mongodb

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/config"
	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type MongoDB struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func New(ctx context.Context, cfg *config.Config) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	opts := clientOpts(cfg)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to send ping: %w", err)
	}

	coll := client.Database(cfg.ConfigMongoDB.DBName).Collection(cfg.ConfigMongoDB.CollName)

	if err := createIndex(ctx, coll); err != nil {
		return nil, fmt.Errorf("failed to create index: %w", err)
	}

	return &MongoDB{
		client: client,
		coll:   coll,
	}, nil
}

func (m *MongoDB) Insert(ctx context.Context, weapons []*types.Weapon) error {
	log := logger.FromContext(ctx, logger.Storage)

	res, err := m.coll.InsertMany(ctx, weapons)
	if err != nil {
		log.Error("database failed",
			zap.Error(err),
			zap.String("operation", "InsertMany"),
		)
		return fmt.Errorf("failed to insert documents: %w", err)
	}

	log.Debug("Insert complited",
		zap.Int("total documents inserted", len(res.InsertedIDs)),
	)

	return nil
}

func (m *MongoDB) WeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	log := logger.FromContext(ctx, logger.Storage)

	filter := bson.M{"category": category}

	cursor, err := m.coll.Find(ctx, filter)
	if err != nil {
		log.Error("database failed",
			zap.Error(err),
			zap.String("operation", "Find"),
		)
		return nil, fmt.Errorf("failed to find documents: %w", err)
	}
	defer cursor.Close(ctx)

	var weapons []*types.Weapon

	if err := cursor.All(ctx, &weapons); err != nil {
		log.Error("database failed",
			zap.Error(err),
			zap.String("operation", "Decoding"),
		)
		return nil, fmt.Errorf("failed to decode document: %w", err)
	}

	if err := cursor.Err(); err != nil {
		log.Error("iteration error",
			zap.Error(err),
		)
		return nil, fmt.Errorf("last cursor error: %w", err)
	}

	log.Debug("WeaponsByCategory complited",
		zap.Any("filter", filter),
		zap.Int("total documents returned", len(weapons)),
	)

	return weapons, nil
}

func (m *MongoDB) Search(ctx context.Context, name string) (map[string]string, error) {
	log := logger.FromContext(ctx, logger.Storage)

	filter := bson.M{
		"name": bson.M{
			"$regex":   "^" + regexp.QuoteMeta(name),
			"$options": "i",
		},
	}

	cursor, err := m.coll.Find(ctx, filter)
	if err != nil {
		log.Error("database failed",
			zap.Error(err),
			zap.String("operation", "Find"),
		)
		return nil, fmt.Errorf("failed to find documents: %w", err)
	}
	defer cursor.Close(ctx)

	result := make(map[string]string, 20)

	for cursor.Next(ctx) {
		var searchResult types.SearchResult
		if err := cursor.Decode(&searchResult); err != nil {
			log.Error("failed to decode document", zap.Error(err))
			continue
		}

		result[searchResult.Name] = searchResult.Category
	}

	if err := cursor.Err(); err != nil {
		log.Error("iteration error",
			zap.Error(err),
		)
		return nil, fmt.Errorf("last cursor error: %w", err)
	}

	return result, nil
}

func (m *MongoDB) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
