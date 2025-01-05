package usersession

import (
	"context"
	"log"
	"net/http"

	"github.com/pohanson/cvwo-forum/internal/model"
)

type ctxKey string

const userCtxKey ctxKey = "user"

func WithUserCtx(ctx context.Context, user model.User) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

func GetUserFromCtx(ctx context.Context) (model.User, bool) {
	user, ok := ctx.Value(userCtxKey).(model.User)
	if !ok {
		log.Println("Error getting user from context", ctx.Value(userCtxKey))
	}
	return user, ok
}
func SessionUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := GetUserFromReq(r)
		if ok {
			ctx := WithUserCtx(r.Context(), user)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
