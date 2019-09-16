package login

import (
	"net/http"
	"sample/app/shared/handler"

	"github.com/gorilla/sessions"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

type templateData struct {
	Mess  string
	Email string
}

var (
	key   = []byte("secret-key")
	store = sessions.NewCookieStore(key)
)

// Login page
func (h *HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	message, _ := session.Values["mess_session"].(string)
	email, _ := session.Values["email"].(string)
	if (message == "" && email == "") || (message == "email or password invalid") {
		err := h.ResponseHTML(w, r, "login/login", templateData{
			Mess:  message,
			Email: email,
		})
		if err != nil {
			_ = h.StatusServerError(w, r)
		}
	} else {
		http.Redirect(w, r, "/logout", 302)
	}
}

//HandleLogin ...
func (h *HTTPHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("pwd")
	// fmt.Println(email)
	// fmt.Println(password)
	if email == "hathaymuadep@gmail.com" && password == "quanhen121" {
		message := "done"
		session, _ := store.Get(r, "user")
		session.Values["mess_session"] = message
		session.Values["email"] = email
		session.Save(r, w)
	} else {
		message := "email or password invalid"
		session, _ := store.Get(r, "user")
		session.Values["mess_session"] = message
		session.Save(r, w)
	}

	http.Redirect(w, r, "/login", 302)
}

// NewLoginHTTPHandler responses new HTTPHandler instance.
func NewLoginHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
