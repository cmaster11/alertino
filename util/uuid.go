package util

import (
	"crypto/md5"
	"fmt"

	"github.com/intelekshual/go-simpleflake"
)

func NewSeqMD5Id() (string, error) {
	flake, err := simpleflake.New()
	if err != nil {
		return "", err
	}

	flakeStr := fmt.Sprintf("%d", flake)
	md5Bytes := md5.Sum([]byte(flakeStr))

	return fmt.Sprintf("%x", md5Bytes), nil
}
