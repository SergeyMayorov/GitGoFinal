package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(h Handler) http.Handler {
	// роутер
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

	router.Get("/books", h.GetAllBooks)
	router.Get("/books/{id}", h.GetBook)
	router.Put("/books/{id}", h.UpdBook)
	router.Post("/books", h.InsBook)
	router.Delete("/books/{id}", h.DelBook)
	router.Get("/authors", h.GetAllAuthors)
	router.Get("/authors/{id}", h.GetAuthor)
	router.Put("/authors/{id}", h.UpdAuthor)
	router.Post("/authors", h.InsAuthor)
	router.Delete("/authors/{id}", h.DelAuthor)
	router.Put("/books/{id_book}/authors/{id_author}", h.UpdAuthorBook)
	return router
}
