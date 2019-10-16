package register

import (
	"net/http"
	"sample/app/shared/handler"

	"github.com/gorilla/sessions"
)

type templateData struct {
	Mess  string
	Email string
}

var (
	key   = []byte("secret-key")
	store = sessions.NewCookieStore(key)
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

var message string

// Register page
func (h *HTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	session.Options.MaxAge = -1
	session.Save(r, w)
	err := h.ResponseHTML(w, r, "login/register", templateData{
		Mess: message,
	})
	if err != nil {
		_ = h.StatusServerError(w, r)
	}
	message = ""
}

//HandlerRegister ...
func (h *HTTPHandler) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	message = ""
	email := r.FormValue("email")
	password := r.FormValue("pwd")
	passwordCf := r.FormValue("pwd_cf")
	if password != passwordCf {
		message = "password confirm not valid"
		http.Redirect(w, r, "/register", 302)
	} else {
		UserInfo := handler.ReturnStruct(email, password)
		check := handler.CheckEmail(UserInfo)
		if check == false {
			message = "you can login now"
			handler.InsertUser(UserInfo)
		} else {
			message = "account exited"
		}
		http.Redirect(w, r, "/register", 302)
	}

}

// NewRegisterHTTPHandler responses new HTTPHandler instance.
func NewRegisterHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
