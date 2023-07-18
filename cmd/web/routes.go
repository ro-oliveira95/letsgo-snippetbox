package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/ro-oliveira95/letsgo-snippetbox/ui"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// Overwrite default httprouter's NotFound handler
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	// Middleware chains
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)
	protected := dynamic.Append(app.requireAuthentication)

	// Routes definition
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignupView))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignup))
	router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLoginView))
	router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLogin))

	router.Handler(http.MethodGet, "/snippet/create", protected.ThenFunc(app.snippetCreateView))
	router.Handler(http.MethodPost, "/snippet/create", protected.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogout))

	return standard.Then(router)
}
