package model

import (
	"encoding/json"
	"io"
)

type Category struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CategoryList struct {
	CategoryList []Category `json:"categoryList"`
}

func (c *CategoryList) DecodeFromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(c)
}

func (c *Category) DecodeFromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&c)
}
