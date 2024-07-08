package repository

import (
	"github.com/SergeyMayorov/GitGoFinal/pkg/models"
)

type DBInterface interface {
	// ----------------- Работа с книгами -------------------------
	GetListBooks() ([]*models.Book, error)
	GetBookById(id int) (*models.Book, error)
	InsBookById(book models.Book) (int, error)
	UpdBookById(book models.Book) error
	DelBookById(id int) error
	// ----------------- Работа с авторами ----------------------
	GetListAuthors() ([]*models.Author, error)
	GetAuthorById(id int) (*models.Author, error)
	InsAuthorById(author models.Author) (int, error)
	UpdAuthorById(author models.Author) error
	DelAuthorById(id int) error
	// ----------------- Работа с авторами и книгами -------------------------
	UpdAuthorBook(author models.Author, book models.Book) error
}
