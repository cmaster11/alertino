package platform

const IndexFieldExpiresAt = "expiresAt"

type DbModel struct {
	Key string `json:"_key"`
}

type ExpiringDbModel struct {
	DbModel

	ExpiresAt int64 `json:"expiresAt"`
}
