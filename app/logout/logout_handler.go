package logout

import (
	"net/http"
	"sample/app/shared/handler"

	"github.com/gorilla/sessions"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

var (
	key   = []byte("secret-key")
	store = sessions.NewCookieStore(key)
)

type templateData struct {
	Email string
	Mess  string
}

// Logout get session
func (h *HTTPHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	email, _ := session.Values["email"].(string)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := h.ResponseHTML(w, r, "logout/logout", templateData{
			Email: email,
		})
		if err != nil {
			_ = h.StatusServerError(w, r)
		}
	}
}

//HandlerLogout delete session
func (h *HTTPHandler) HandlerLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}

// NewLogoutHTTPHandler responses new HTTPHandler instance.
func NewLogoutHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
