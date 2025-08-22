package mongodb

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/config"
	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

const (
	FieldWeaponID        = "id"
	FieldWeaponsCategory = "category"
	FieldWeaponName      = "name"
	FieldVersionID       = "_id"
	CurrentVersion       = "current_version"
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

	return &MongoDB{
		client: client,
		coll:   coll,
	}, nil
}

func (m *MongoDB) UpsertWeapons(ctx context.Context, weapons []*types.Weapon) error {
	log := logger.FromContext(ctx, logger.Storage)

	models := make([]mongo.WriteModel, 0, len(weapons))

	for _, weapon := range weapons {
		if weapon.ID == "" {
			weapon.ID = generateWeaponID(weapon)
		}

		filter := bson.M{FieldWeaponID: weapon.ID}
		update := updateWeapon(weapon)

		model := mongo.NewUpdateOneModel()
		model.SetFilter(filter)
		model.SetUpdate(update)
		model.SetUpsert(true)

		models = append(models, model)
	}

	res, err := m.coll.BulkWrite(ctx, models)
	if err != nil {
		log.Error("BulkWrite error",
			zap.Error(err),
		)
		return fmt.Errorf("failed to upsert documents: %w", err)
	}

	log.Debug("UpsertWeapons complited",
		zap.Int("matched count", int(res.MatchedCount)),
		zap.Int("upserted count", int(res.UpsertedCount)),
		zap.Int("modified count", int(res.ModifiedCount)),
	)

	return nil
}

func (m *MongoDB) WeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	log := logger.FromContext(ctx, logger.Storage)

	filter := bson.M{FieldWeaponsCategory: category}

	cursor, err := m.coll.Find(ctx, filter)
	if err != nil {
		log.Error("Find error",
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to find documents: %w", err)
	}
	defer cursor.Close(ctx)

	var weapons []*types.Weapon

	if err := cursor.All(ctx, &weapons); err != nil {
		log.Error("Decode error",
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to decode document: %w", err)
	}

	if err := cursor.Err(); err != nil {
		log.Error("Cursor error",
			zap.Error(err),
		)
		return nil, fmt.Errorf("last cursor error: %w", err)
	}

	log.Debug("WeaponsByCategory complited",
		zap.Int("total documents found", len(weapons)),
	)

	return weapons, nil
}

func (m *MongoDB) WeaponsByName(ctx context.Context, query string) ([]types.SearchResult, error) {
	log := logger.FromContext(ctx, logger.Storage)

	filter := bson.M{
		FieldWeaponName: bson.M{
			"$regex":   regexp.QuoteMeta(query),
			"$options": "i",
		},
	}

	cursor, err := m.coll.Find(ctx, filter)
	if err != nil {
		log.Error("Find error",
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to find documents: %w", err)
	}
	defer cursor.Close(ctx)

	var results []types.SearchResult

	if err := cursor.All(ctx, &results); err != nil {
		log.Error("Decode error",
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to decode document: %w", err)
	}

	if err := cursor.Err(); err != nil {
		log.Error("Cursor error",
			zap.Error(err),
		)
		return nil, fmt.Errorf("last cursor error: %w", err)
	}

	log.Debug("WeaponsByName complited",
		zap.Int("total documents found", len(results)),
	)

	return results, nil
}

func (m *MongoDB) Version(ctx context.Context) (types.LastChange, error) {
	log := logger.FromContext(ctx, logger.Storage)

	filter := bson.M{FieldVersionID: CurrentVersion}

	var version types.LastChange

	err := m.coll.FindOne(ctx, filter).Decode(&version)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Error("Version not found",
				zap.Error(err),
			)
			return types.LastChange{}, fmt.Errorf("version not found: %w", err)
		}
		log.Error("Decode error",
			zap.Error(err),
		)
		return types.LastChange{}, fmt.Errorf("failed to find document: %w", err)
	}

	log.Debug("Version complited",
		zap.Any("found document", version),
	)

	return version, nil
}

func (m *MongoDB) UpsertVersion(ctx context.Context, version types.VersionInfo) error {
	log := logger.FromContext(ctx, logger.Storage)

	opts := options.UpdateOne().SetUpsert(true)

	filter := bson.M{FieldVersionID: CurrentVersion}
	update := updateVersion(version)

	res, err := m.coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Error("Update error",
			zap.Error(err),
		)
		return fmt.Errorf("failed to update document: %w", err)
	}

	log.Debug("UpsetVersion complited",
		zap.Int("matched count", int(res.MatchedCount)),
		zap.Int("upserted count", int(res.UpsertedCount)),
		zap.Int("modified count", int(res.ModifiedCount)),
	)

	return nil
}

func (m *MongoDB) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
