package mongodb

import (
	"context"
	"fmt"
	"testing"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	c "github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var notExistentCategories = []string{"", "aam-rear", "aam_ir_rear_aspect", "aam ir all aspect"}

func setupMongo(ctx context.Context) (*mongo.Client, error) {
	container, err := c.Run(ctx, "mongo:8.0")
	if err != nil {
		return nil, err
	}

	connStr, err := container.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}

	return mongo.Connect(options.Client().ApplyURI(connStr))
}

func TestMongoDB_Insert(t *testing.T) {
	ctx := context.Background()

	client, err := setupMongo(ctx)
	require.NoError(t, err)

	coll := client.Database("test").Collection("test_weapons")

	db := &MongoDB{
		client: client,
		coll:   coll,
	}

	weapons := []*types.Weapon{
		{Category: "atgm-ir", Name: "QN502C"},
		{Category: "atgm-ir", Name: "Spike LR2"},
		{Category: "atgm-losbr", Name: "ACRA"},
	}

	t.Run("Success Insert inserts slice of weapons", func(t *testing.T) {
		err := db.Insert(ctx, weapons)
		require.NoError(t, err)

		count, err := coll.CountDocuments(ctx, bson.M{})
		require.NoError(t, err)
		assert.Equal(t, len(weapons), int(count))

		var res types.Weapon
		err = coll.FindOne(ctx, bson.M{"category": "atgm-losbr"}).Decode(&res)
		require.NoError(t, err)
		assert.Equal(t, "ACRA", res.Name)
	})

	t.Cleanup(func() {
		coll.Drop(ctx)
		client.Disconnect(ctx)
	})
}

func TestMongoDB_WeaponsByCategory(t *testing.T) {
	ctx := context.Background()

	client, err := setupMongo(ctx)
	require.NoError(t, err)

	coll := client.Database("test").Collection("test_weapons")

	_, err = coll.InsertMany(ctx, []interface{}{
		bson.M{"category": "aam-ir-rear-aspect", "name": "AIM-9B"},
		bson.M{"category": "aam-ir-rear-aspect", "name": "A-91"},
		bson.M{"category": "aam-ir-all-aspect", "name": "AIM-9L"},
		bson.M{"category": "gbu-tv", "name": "GBU-8/B"},
	})
	require.NoError(t, err)

	db := &MongoDB{
		client: client,
		coll:   coll,
	}

	t.Run("Success WeaponsByCategory returns slice of weapons", func(t *testing.T) {
		weapons, err := db.WeaponsByCategory(ctx, "aam-ir-rear-aspect")
		require.NoError(t, err)
		assert.Len(t, weapons, 2)
		assert.Equal(t, "aam-ir-rear-aspect", weapons[0].Category)
		assert.Equal(t, "AIM-9B", weapons[0].Name)
	})

	for _, category := range notExistentCategories {
		t.Run(fmt.Sprintf("Non-existent category %s returns empty slice of weapons", category), func(t *testing.T) {
			weapons, err := db.WeaponsByCategory(ctx, category)
			require.NoError(t, err)
			assert.Empty(t, weapons)
		})
	}

	t.Cleanup(func() {
		defer client.Disconnect(ctx)
	})
}
