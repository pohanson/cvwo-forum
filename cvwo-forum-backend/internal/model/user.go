package model

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"

	"github.com/pohanson/cvwo-forum/internal/utils"
)

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     int    `json:"role"`
}

func (u *User) DecodeFromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&u)

}

func (u *User) UnmarshalJSON(data []byte) error {
	all := struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Role     int    `json:"role"`
	}{}

	if err := json.Unmarshal(data, &all); err != nil {
		return err
	}

	u.Username = all.Username
	u.Name = all.Name
	u.Role = all.Role

	return u.ValidateAll()
}

func (u *User) ValidateAll() error {
	return errors.Join(
		u.ValidateRequiredField([]string{"Username", "Name"}),
		u.ValidateMaxLength(80, []string{"Username", "Name"}),
	)
}

func (u *User) ValidateRequiredField(requiredFields []string) error {
	emptyRequiredFields := utils.Filter(requiredFields, func(s string) bool {
		return reflect.ValueOf(u).Elem().FieldByName(s).String() == ""
	})
	if len(emptyRequiredFields) != 0 {
		return &MissingFieldErr{emptyRequiredFields}
	}
	return nil
}

func (u *User) ValidateMaxLength(maxLength int, lengthCheckFields []string) error {
	exceed := utils.Filter(lengthCheckFields, func(s string) bool { return reflect.ValueOf(*u).FieldByName(s).Len() > maxLength })
	if len(exceed) != 0 {
		return &MaxLengthViolationErr{maxLength, exceed}
	}
	return nil
}
