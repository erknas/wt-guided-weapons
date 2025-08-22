package mongodb

import (
	"context"
	"testing"

	"github.com/erknas/wt-guided-weapons/internal/config"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

const configPath = "../../../configs/mongodb-test-config.yaml"

func TestMongoDB_UpsertWeapons(t *testing.T) {
	ctx := context.Background()

	cfg := config.MustLoad(configPath)

	db, err := New(ctx, cfg)
	require.NoError(t, err)
	defer db.Close(ctx)

	_, err = db.coll.DeleteMany(ctx, bson.M{})
	require.NoError(t, err)

	t.Run("insert new weapons", func(t *testing.T) {
		weapons := []*types.Weapon{
			{ID: "1a", Name: "AIM-9L", Category: "aam-ir-all-aspect"},
			{ID: "2b", Name: "AIM-54", Category: "aam-arh"},
		}

		err := db.UpsertWeapons(ctx, weapons)
		require.NoError(t, err)

		count, err := db.coll.CountDocuments(ctx, bson.M{})
		require.NoError(t, err)
		assert.Equal(t, int64(2), count)

		_, err = db.coll.DeleteMany(ctx, bson.M{})
		require.NoError(t, err)
	})

	t.Run("update weapons", func(t *testing.T) {
		weapons := []*types.Weapon{
			{ID: "1", Name: "AIM-9L", Category: "aam-ir-all-aspect", Mass: "280"},
			{ID: "2", Name: "AIM-54", Category: "aam-arh", MassAtEndOfBoosterBurn: "450"},
		}

		err := db.UpsertWeapons(ctx, weapons)
		require.NoError(t, err)

		count, err := db.coll.CountDocuments(ctx, bson.M{})
		require.NoError(t, err)
		assert.Equal(t, int64(2), count)

		weapons[0].Mass = "270"
		weapons[1].MassAtEndOfBoosterBurn = "400"

		err = db.UpsertWeapons(ctx, weapons)
		require.NoError(t, err)

		count, err = db.coll.CountDocuments(ctx, bson.M{})
		require.NoError(t, err)
		assert.Equal(t, int64(2), count)

		var weapon types.Weapon

		err = db.coll.FindOne(ctx, bson.M{"id": "1"}).Decode(&weapon)
		require.NoError(t, err)
		assert.Equal(t, "270", weapon.Mass)

		err = db.coll.FindOne(ctx, bson.M{"id": "2"}).Decode(&weapon)
		require.NoError(t, err)
		assert.Equal(t, "400", weapon.MassAtEndOfBoosterBurn)

		_, err = db.coll.DeleteMany(ctx, bson.M{})
		require.NoError(t, err)
	})
}

func TestMongoDB_WeaponsByCategory(t *testing.T) {
	ctx := context.Background()

	cfg := config.MustLoad(configPath)

	db, err := New(ctx, cfg)
	require.NoError(t, err)
	defer db.Close(ctx)

	_, err = db.coll.DeleteMany(ctx, bson.M{})
	require.NoError(t, err)

	t.Run("find weapons by category", func(t *testing.T) {
		weapons := []*types.Weapon{
			{ID: "1", Name: "AIM-9L", Category: "aam-ir-all-aspect"},
			{ID: "2", Name: "AIM-9M", Category: "aam-ir-all-aspect"},
			{ID: "3", Name: "AAM-3", Category: "aam-ir-all-aspect"},
			{ID: "4", Name: "AIM-54", Category: "aam-arh"},
		}

		err := db.UpsertWeapons(ctx, weapons)
		require.NoError(t, err)

		count, err := db.coll.CountDocuments(ctx, bson.M{})
		require.NoError(t, err)
		assert.Equal(t, int64(4), count)

		category := "aam-ir-all-aspect"

		results, err := db.WeaponsByCategory(ctx, category)
		require.NoError(t, err)
		assert.Equal(t, 3, len(results))
		assert.Equal(t, weapons[0].Name, results[0].Name)
		assert.Equal(t, weapons[1].Name, results[1].Name)
		assert.Equal(t, weapons[2].Name, results[2].Name)

		_, err = db.coll.DeleteMany(ctx, bson.M{})
		require.NoError(t, err)
	})
}

func TestMongoDB_WeaponsByName(t *testing.T) {
	ctx := context.Background()

	cfg := config.MustLoad(configPath)

	db, err := New(ctx, cfg)
	require.NoError(t, err)
	defer db.Close(ctx)

	_, err = db.coll.DeleteMany(ctx, bson.M{})
	require.NoError(t, err)

	t.Run("find weapons by name", func(t *testing.T) {
		weapons := []*types.Weapon{
			{ID: "1", Name: "AIM-9L", Category: "aam-ir-all-aspect"},
			{ID: "2", Name: "AIM-9M", Category: "aam-ir-all-aspect"},
			{ID: "3", Name: "AAM-3", Category: "aam-ir-all-aspect"},
			{ID: "4", Name: "AIM-54", Category: "aam-arh"},
		}

		err := db.UpsertWeapons(ctx, weapons)
		require.NoError(t, err)

		count, err := db.coll.CountDocuments(ctx, bson.M{})
		require.NoError(t, err)
		assert.Equal(t, int64(4), count)

		query := "aim"

		results, err := db.WeaponsByName(ctx, query)
		require.NoError(t, err)
		assert.Equal(t, 3, len(results))

		_, err = db.coll.DeleteMany(ctx, bson.M{})
		require.NoError(t, err)
	})

	t.Run("find weapons by name empty results", func(t *testing.T) {
		weapons := []*types.Weapon{
			{ID: "1", Name: "AIM-9L", Category: "aam-ir-all-aspect"},
			{ID: "2", Name: "AIM-9M", Category: "aam-ir-all-aspect"},
			{ID: "3", Name: "AAM-3", Category: "aam-ir-all-aspect"},
			{ID: "4", Name: "AIM-54", Category: "aam-arh"},
		}

		err := db.UpsertWeapons(ctx, weapons)
		require.NoError(t, err)

		count, err := db.coll.CountDocuments(ctx, bson.M{})
		require.NoError(t, err)
		assert.Equal(t, int64(4), count)

		query := "abfa1230"

		results, err := db.WeaponsByName(ctx, query)
		require.NoError(t, err)
		assert.Empty(t, results)

		_, err = db.coll.DeleteMany(ctx, bson.M{})
		require.NoError(t, err)
	})
}

func TestMongoDB_Version(t *testing.T) {
	ctx := context.Background()

	cfg := config.MustLoad(configPath)

	db, err := New(ctx, cfg)
	require.NoError(t, err)
	defer db.Close(ctx)

	_, err = db.coll.DeleteMany(ctx, bson.M{})
	require.NoError(t, err)

	t.Run("get version", func(t *testing.T) {
		version := types.VersionInfo{Version: "2.46.0.114"}

		err := db.UpsertVersion(ctx, version)
		require.NoError(t, err)

		results, err := db.Version(ctx)
		require.NoError(t, err)
		assert.NotEmpty(t, results)

		_, err = db.coll.DeleteMany(ctx, bson.M{})
		require.NoError(t, err)
	})

	t.Run("version not found", func(t *testing.T) {
		results, err := db.Version(ctx)
		require.Error(t, err)
		assert.Empty(t, results)
		assert.Contains(t, err.Error(), "version not found")
	})
}

func TestMongoDB_UpsertVersion(t *testing.T) {
	ctx := context.Background()

	cfg := config.MustLoad(configPath)

	db, err := New(ctx, cfg)
	require.NoError(t, err)
	defer db.Close(ctx)

	_, err = db.coll.DeleteMany(ctx, bson.M{})
	require.NoError(t, err)

	t.Run("insert version", func(t *testing.T) {
		version := types.VersionInfo{Version: "2.46.0.114"}

		err := db.UpsertVersion(ctx, version)
		require.NoError(t, err)

		results, err := db.Version(ctx)
		require.NoError(t, err)
		assert.NotEmpty(t, results)

		assert.Equal(t, results.Version, version)

		_, err = db.coll.DeleteMany(ctx, bson.M{})
		require.NoError(t, err)
	})
}
