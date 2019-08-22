package router

import (
	"net/http"

	"sample/app/hello"
	"sample/app/infrastructure"
	"sample/app/login"
	"sample/app/shared/handler"
	middlewareAuth "sample/app/shared/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router is application struct hold Mux and db connection
type Router struct {
	Mux                *chi.Mux
	SQLHandler         *infrastructure.SQL
	CacheHandler       *infrastructure.Cache
	LoggerHandler      *infrastructure.Logger
	TranslationHandler *infrastructure.Translation
}

// InitializeRouter initializes Mux and middleware
func (r *Router) InitializeRouter() {
	r.Mux.Use(middleware.RequestID)
	r.Mux.Use(middleware.RealIP)
	// Custom middleware(Translation)
	// r.Mux.Use(r.TranslationHandler.Middleware.Middleware)
	// // Custom middleware(Logger)
	// r.Mux.Use(mMiddleware.Logger(r.LoggerHandler))

}

// SetupHandler set database and redis and usecase.
func (r *Router) SetupHandler() {
	// error handler set.
	eh := handler.NewHTTPErrorHandler(r.LoggerHandler.Log)
	r.Mux.NotFound(eh.StatusNotFound)
	r.Mux.MethodNotAllowed(eh.StatusMethodNotAllowed)

	r.Mux.Method(http.MethodGet, "/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Mux.HandleFunc("/terms-of-use", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/terms-of-use.html")
	})

	ah := handler.NewApplicationHTTPHandler(r.LoggerHandler.Log)

	// base set.
	// bh := handler.NewBaseHTTPHandler(r.LoggerHandler.Log)
	// // base set.
	// br := repository.NewBaseRepository(r.LoggerHandler.Log)
	// // base set.
	// bu := usecase.NewBaseUsecase(r.LoggerHandler.Log)

	// uh := user.NewHTTPHandler(br, bu, bh, r.SQLHandler, r.CacheHandler)

	hw := hello.NewHTTPHandler(ah)
	lg := login.NewHTTPHandler(ah)

	r.Mux.Route("/", func(cr chi.Router) {
		cr.Route("/hello", func(cr chi.Router) {
			cr.Use(middlewareAuth.VerifyAuth)
			cr.Get("/", hw.HelloWorld)
		})
		cr.Get("/login", lg.Login)
		cr.Post("/login", lg.LoginHandler)
		cr.Get("/logout", lg.LogoutHandler)
	})
}
