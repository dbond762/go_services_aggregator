package messages

import (
	"net/http"
)

const cookieName = "message"

type Message struct {
	Value string
	Valid bool
}

func Set(w http.ResponseWriter, message string) {
	cookie := &http.Cookie{
		Name:  "message",
		Value: message,
	}

	http.SetCookie(w, cookie)
}

func Flush(w http.ResponseWriter, r *http.Request) Message {
	cookie, err := r.Cookie(cookieName)
	if err == http.ErrNoCookie || cookie == nil {
		return Message{}
	}

	message := cookie.Value

	cookie.MaxAge = -1
	http.SetCookie(w, cookie)

	return Message{
		Value: message,
		Valid: true,
	}
}
