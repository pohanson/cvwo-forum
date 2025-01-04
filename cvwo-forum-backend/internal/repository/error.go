package repository

import (
	"fmt"
)

type UniqueViolationErr struct {
	field string
}

func (e UniqueViolationErr) Error() string {
	return fmt.Sprintf("%s is taken!", e.field)
}
