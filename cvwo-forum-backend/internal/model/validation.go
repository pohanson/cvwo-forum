package model

import (
	"reflect"

	"github.com/pohanson/cvwo-forum/internal/utils"
)

func ValidateRequiredField(u interface{}, requiredFields []string) error {
	emptyRequiredFields := utils.Filter(requiredFields, func(s string) bool {
		return reflect.ValueOf(u).Elem().FieldByName(s).String() == ""
	})
	if len(emptyRequiredFields) != 0 {
		return &MissingFieldErr{emptyRequiredFields}
	}
	return nil
}
