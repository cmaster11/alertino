package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"regexp"
)

var Validate = validator.New()

type Validable struct {
}

func (v *Validable) Validate() error {
	if err := Validate.Struct(v); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}

func (v *Validable) PanicIfInvalid() {
	if v == nil {
		logrus.Panic("nil struct")
	}

	if err := v.Validate(); err != nil {
		logrus.Panic(err)
	}
}

func init() {
	regexUrlFriendly := regexp.MustCompile(`^[\w-]+$`)

	PanicIfError(Validate.RegisterValidation("urlfriendly", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		if !regexUrlFriendly.MatchString(value) {
			return false
		}
		return true
	}, true))
}
