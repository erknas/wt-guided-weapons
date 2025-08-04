package mongodb

import (
	"fmt"
	"net/url"

	"github.com/erknas/wt-guided-weapons/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func clientOpts(cfg *config.Config) *options.ClientOptions {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		url.QueryEscape(cfg.ConfigMongoDB.Username),
		url.QueryEscape(cfg.ConfigMongoDB.Password),
		cfg.ConfigMongoDB.Host,
		cfg.ConfigMongoDB.Port,
	)

	opts := options.Client().
		ApplyURI(uri).
		SetConnectTimeout(cfg.ConfigMongoDB.ConnectTimeout).
		SetServerSelectionTimeout(cfg.ConfigMongoDB.SelectTimeout)

	return opts
}
