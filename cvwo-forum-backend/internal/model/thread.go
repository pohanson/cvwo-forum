package model

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

type ThreadType int

const (
	QuestionThreadType ThreadType = 1
	PostThreadType     ThreadType = 2
	ReplyThreadType    ThreadType = 3
)

type Thread struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Type      ThreadType `json:"type"`
	CreatedBy int        `json:"createdBy"`
	CreatedOn time.Time  `json:"createdOn"`
	EditedOn  time.Time  `json:"editedOn"`
}

func (t *Thread) DecodeFromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&t)
}

func (t *Thread) ValidateAll() error {
	return errors.Join(t.ValidateRequiredField())
}
func (t *Thread) ValidateRequiredField() error {
	switch t.Type {
	case PostThreadType:
		return ValidateRequiredField(t, []string{"Title", "Type", "Content", "CreatedBy"})
	case QuestionThreadType:
		fallthrough
	case ReplyThreadType:
		return ValidateRequiredField(t, []string{"Title", "Type", "CreatedBy"})
	default:
		return &InvalidFK{"Thread Type"}
	}
}
func (t *Thread) ValidateThreadType() error {
	if t.Type >= 1 || t.Type <= 3 {
		return nil
	} else {
		return &InvalidFK{"Thread Type"}
	}
}
