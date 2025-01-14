package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/pohanson/cvwo-forum/internal/model"
	"github.com/pohanson/cvwo-forum/internal/repository"
	"github.com/pohanson/cvwo-forum/internal/usersession"
)

type customPostJson struct {
	categoryList *model.CategoryList
	thread       *model.Thread
}

func (c *customPostJson) DecodeFromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&c)
}

func (c *customPostJson) UnmarshalJSON(data []byte) error {
	return errors.Join(json.Unmarshal(data, &c.thread),
		json.Unmarshal(data, &c.categoryList))

}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c := customPostJson{}
	if err := c.DecodeFromJson(r.Body); err != nil {
		log.Println("Error decoding json body:", err)
		writeDecodingError(err, w)
	}

	thread := c.thread
	categoryList := c.categoryList

	user, ok := usersession.GetUserFromCtx(ctx)
	if !ok {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}
	thread.Type = model.PostThreadType
	thread.CreatedBy = user.Id
	if err := thread.ValidateAll(); err != nil {
		writeValidationError(err, w)
		log.Println(err)
		return
	}

	repo := repository.GetPgRepo()

	thread, err := repo.Thread.Create(ctx, thread)
	if err != nil {
		log.Println(err)
		writeDbInsertError(err, w)
		return
	}
	errList := make([]error, 0)
	for _, c := range categoryList.CategoryList {
		if c.Id == -1 {
			cat, err := repo.Category.Create(ctx, &c)
			if err == nil {
				errList = append(errList, repo.ThreadCategory.Create(ctx, thread, cat))
			}
		}
	}
	if err := errors.Join(errList...); err != nil {
		log.Println(err)
		http.Error(w, "Unknown Error", http.StatusBadRequest)
	} else {
		jsonResult, _ := json.Marshal(`{"success":true}`)
		w.Write(jsonResult)
	}
}
