package model

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"io"
	"reflect"

	"github.com/pohanson/cvwo-forum/internal/utils"
)

func init() {
	gob.Register(User{})
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     int    `json:"role"`
}

func (u *User) DecodeFromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&u)

}

func (u *User) ValidateAll() error {
	return errors.Join(
		u.ValidateRequiredField(),
		u.ValidateMaxLength(80, []string{"Username", "Name"}),
	)
}

func (u *User) ValidateRequiredField() error {
	return ValidateRequiredField(u, []string{"Username", "Name"})
}

func (u *User) ValidateMaxLength(maxLength int, lengthCheckFields []string) error {
	exceed := utils.Filter(lengthCheckFields, func(s string) bool { return reflect.ValueOf(*u).FieldByName(s).Len() > maxLength })
	if len(exceed) != 0 {
		return &MaxLengthViolationErr{maxLength, exceed}
	}
	return nil
}
