package mongodb

import (
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.mongodb.org/mongo-driver/bson"
)

func updateVersion(version types.VersionInfo) bson.M {
	return bson.M{
		"$set": bson.M{
			"version": version,
		},
	}
}
