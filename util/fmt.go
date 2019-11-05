package util

import (
	"fmt"
)

func Dump(el interface{}) string {
	return fmt.Sprintf("%+v", el)
}
