package usersession

import (
	"log"
	"net/http"
	"os"

	"github.com/golangcollege/sessions"
	"github.com/pohanson/cvwo-forum/internal/model"
)

var session *sessions.Session

const userSesKey string = "user"

func init() {
	var secret = []byte(os.Getenv("SES_SECRET"))
	session = sessions.New(secret)
}

func GetSession() *sessions.Session {
	return session
}
func PutSesUser(r *http.Request, user model.User) {
	session.Put(r, userSesKey, user)
}

func RemoveSesUser(r *http.Request) {
	session.Remove(r, userSesKey)
}

func GetUserFromReq(r *http.Request) (model.User, bool) {
	user, ok := session.Get(r, userSesKey).(model.User)
	if !ok {
		log.Println("Error getting user from request: ", session.Get(r, userSesKey))
	}
	return user, ok
}
