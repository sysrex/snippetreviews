package main

import (
	"net/http"
	"github.com/justinas/alice"
	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

    mux := pat.New()
    mux.Get("/", http.HandlerFunc(app.home))
    mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
    mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
    mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Get("/static/", http.StripPrefix("/static", fileServer))

    // Return the 'standard' middleware chain followed by the servemux.
    return standardMiddleware.Then(mux)
}
