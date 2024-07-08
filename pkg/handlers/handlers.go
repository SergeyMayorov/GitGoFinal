package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/SergeyMayorov/GitGoFinal/pkg/models"
	"github.com/SergeyMayorov/GitGoFinal/pkg/repository"
)

type Application struct {
	Domain  string
	DB      repository.DBInterface
	AppPort string
}

type Handler interface {
	Start(h http.Handler) error

	GetAllBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	InsBook(w http.ResponseWriter, r *http.Request)
	UpdBook(w http.ResponseWriter, r *http.Request)
	DelBook(w http.ResponseWriter, r *http.Request)

	GetAllAuthors(w http.ResponseWriter, r *http.Request)
	GetAuthor(w http.ResponseWriter, r *http.Request)
	UpdAuthor(w http.ResponseWriter, r *http.Request)
	InsAuthor(w http.ResponseWriter, r *http.Request)
	DelAuthor(w http.ResponseWriter, r *http.Request)
	UpdAuthorBook(w http.ResponseWriter, r *http.Request)
}

func (app *Application) Start(h http.Handler) error {
	err := http.ListenAndServe(app.AppPort, h)
	if err != nil {
		return err
	}
	log.Println("Succ start on port %s", app.AppPort)
	return nil
}

func New(db repository.DBInterface, port string) Handler {
	return &Application{
		DB:      db,
		AppPort: port,
	}
}

func (app *Application) GetAllBooks(w http.ResponseWriter, r *http.Request) {

	books, err := app.DB.GetListBooks()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, books)

}

func (app *Application) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
	}

	book, err := app.DB.GetBookById(bookID)
	if err != nil {
		app.errorJSON(w, err)
	}

	_ = app.writeJSON(w, http.StatusOK, book)
}

func (app *Application) InsBook(w http.ResponseWriter, r *http.Request) {
	var bookWithID models.Book

	err := app.readJSON(w, r, &bookWithID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	newID, err := app.DB.InsBookById(bookWithID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Книга с id %d успешно добавлена", newID),
	}

	app.writeJSON(w, http.StatusCreated, resp)

}

func (app *Application) UpdBook(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
	}

	var payload models.Book

	err = app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	book, err := app.DB.GetBookById(ID)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	book.Title = payload.Title
	book.AuthorID = payload.AuthorID
	book.Year = payload.Year
	book.ISBN = payload.ISBN

	err = app.DB.UpdBookById(*book)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Книга c id %d успешно обновлена", ID),
	}

	app.writeJSON(w, http.StatusOK, resp)

}

func (app *Application) DelBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
	}

	_, err = app.DB.GetBookById(ID)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	err = app.DB.DelBookById(ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Книга c id %d успешно удалена", ID),
	}

	app.writeJSON(w, http.StatusOK, resp)

}

// с этой строки и ниже действия с авторами

func (app *Application) GetAllAuthors(w http.ResponseWriter, r *http.Request) {

	books, err := app.DB.GetListAuthors()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, books)

}

func (app *Application) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
	}

	author, err := app.DB.GetAuthorById(authorID)
	if err != nil {
		app.errorJSON(w, err)
	}

	_ = app.writeJSON(w, http.StatusOK, author)
}

func (app *Application) InsAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	err := app.readJSON(w, r, &author)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	newID, err := app.DB.InsAuthorById(author)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Автор с id %d успешно добавлен", newID),
	}

	app.writeJSON(w, http.StatusCreated, resp)

}

func (app *Application) UpdAuthor(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
	}

	var payload models.Author

	err = app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	author, err := app.DB.GetAuthorById(ID)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	author.Name = payload.Name
	author.Sirname = payload.Sirname
	author.Biography = payload.Biography
	author.Birthday = payload.Birthday

	err = app.DB.UpdAuthorById(*author)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Автор с id %d успешно обновлен", ID),
	}

	app.writeJSON(w, http.StatusOK, resp)

}

func (app *Application) DelAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		app.errorJSON(w, err)
	}

	_, err = app.DB.GetAuthorById(ID)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	err = app.DB.DelAuthorById(ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: fmt.Sprintf("Автор с id %d успешно удален", ID),
	}

	app.writeJSON(w, http.StatusOK, resp)

}

func (app *Application) UpdAuthorBook(w http.ResponseWriter, r *http.Request) {
	id_book := chi.URLParam(r, "id_book")
	ID_book, err := strconv.Atoi(id_book)
	if err != nil {
		app.errorJSON(w, err)
	}

	id_author := chi.URLParam(r, "id_author")
	ID_author, err := strconv.Atoi(id_author)
	if err != nil {
		app.errorJSON(w, err)
	}

	var payload models.RelAuthorBook

	err = app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	author, err := app.DB.GetAuthorById(ID_author)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	book, err := app.DB.GetBookById(ID_book)
	if err != nil {
		app.errorJSON(w, err, 404)
		return
	}

	author.Name = payload.Author.Name
	author.Sirname = payload.Author.Sirname
	author.Biography = payload.Author.Biography
	author.Birthday = payload.Author.Birthday
	book.Title = payload.Book.Title
	book.AuthorID = ID_author
	book.Year = payload.Book.Year
	book.ISBN = payload.Book.ISBN

	err = app.DB.UpdAuthorBook(*author, *book)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponce{
		Error:   false,
		Message: "Автор и книга успешно обновлены",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
