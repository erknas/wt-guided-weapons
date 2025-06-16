package storage

import (
	"context"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/config"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	_, err := m.coll.InsertMany(ctx, weapons)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDB) Update(ctx context.Context, weapons []*types.Weapon) error {
	return nil
}

func (m *MongoDB) WeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	return nil, nil
}

func (m *MongoDB) Weapons(ctx context.Context) ([]*types.Weapon, error) {
	cursor, err := m.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var weapons []*types.Weapon
	if err := cursor.All(ctx, &weapons); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return weapons, nil
}

func (m *MongoDB) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
