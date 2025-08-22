package types

type VersionInfo struct {
	Version string `json:"version" bson:"version"`
}

type LastChange struct {
	Version VersionInfo
}
