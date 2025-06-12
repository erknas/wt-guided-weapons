package storage

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/config"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (m *MongoDB) Insert(ctx context.Context, params *types.WeaponParams) error {
	return nil
}

func (m *MongoDB) Update(ctx context.Context, params *types.WeaponParams) error {
	return nil
}

func (m *MongoDB) Provide(ctx context.Context, category string) ([]*types.WeaponParams, error) {
	return nil, nil
}

func (m *MongoDB) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func clientOpts(cfg *config.Config) *options.ClientOptions {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		url.QueryEscape(cfg.ConfigMongoDB.Username),
		url.QueryEscape(cfg.ConfigMongoDB.Password),
		cfg.ConfigMongoDB.Host,
		cfg.ConfigMongoDB.Port,
	)

	opts := options.Client().
		ApplyURI(uri).
		SetConnectTimeout(5 * time.Second).
		SetServerSelectionTimeout(10 * time.Second)

	return opts
}
