package model

import (
	"fmt"
	"strings"
)

type MissingFieldErr struct {
	missingFields []string
}

func (e MissingFieldErr) Error() string {
	return fmt.Sprint("Required Field is Empty: ",
		strings.Join(e.missingFields, ", "))
}

type MaxLengthViolationErr struct {
	maxLength int
	field     []string
}

func (e *MaxLengthViolationErr) Error() string {
	return fmt.Sprintf("%s exceeds the max character length of %d",
		strings.Join(e.field, ", "), e.maxLength)
}

type InvalidFK struct {
	fieldName string
}

func (e *InvalidFK) Error() string {
	return fmt.Sprint(e.fieldName, " has invalid value.")
}
