package util

import "github.com/sirupsen/logrus"

func PanicIfError(err error) {
	if err != nil {
		logrus.Panic(err)
	}
}
