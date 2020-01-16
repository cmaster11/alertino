package auth

import (
	"github.com/cmaster11/alertino/platform"
)

type ModelSession struct {
	platform.ExpiringDbModel
}
