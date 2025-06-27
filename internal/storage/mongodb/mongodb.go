package mongodb

import (
	"context"
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
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	coll := client.Database(cfg.ConfigMongoDB.DBName).Collection(cfg.ConfigMongoDB.CollName)

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
		return nil, err
	}
	defer cursor.Close(ctx)

	var weapons []*types.Weapon

	if err := cursor.All(ctx, &weapons); err != nil {
		log.Error("database failed",
			zap.Error(err),
			zap.String("operation", "Decoding"),
		)
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		log.Error("iteration error",
			zap.Error(err),
		)
		return nil, err
	}

	log.Debug("WeaponsByCategory complited",
		zap.Any("filter", filter),
		zap.Int("total documents returned", len(weapons)),
	)

	return weapons, nil
}

func (m *MongoDB) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
