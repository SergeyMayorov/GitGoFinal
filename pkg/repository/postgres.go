package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/SergeyMayorov/GitGoFinal/pkg/models"
)

type pgDB struct {
	db *sql.DB
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

const dbTimeout = time.Second * 3

func New(cfg Config) DBInterface {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Not connect to DB: %v", err)
	}

	return &pgDB{
		db: db,
	}
}

func (m *pgDB) GetListBooks() ([]*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
select b.id, b.name, a."name" ||' '|| a.sirname as author, b."year" , b.isbn
        from books b
        inner join authors a on b.authorid = a.id
        order by b."name";
		`

	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book

	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Year,
			&book.ISBN,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}

func (m *pgDB) GetBookById(id int) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
select b.id, b.name, a."name" ||' '|| a.sirname as author, b."year" , b.isbn
        from books b
        inner join authors a on b.authorid = a.id
        where b.id =$1;
    `
	row := m.db.QueryRowContext(ctx, query, id)
	var book models.Book
	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Year,
		&book.ISBN,
	)
	if err != nil {
		return nil, err
	}
	return &book, err
}

func (m *pgDB) InsBookById(book models.Book) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        insert into books("name", authorid, year,isbn)
		values($1, $2, $3, $4) returning id
    `

	var newID int

	err := m.db.QueryRowContext(ctx, query,
		book.Title,
		book.AuthorID,
		book.Year,
		book.ISBN,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (m *pgDB) UpdBookById(book models.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        update books
		set "name" = $1, authorid = $2, year = $3, isbn =$4
		where id = $5
    `

	_, err := m.db.ExecContext(ctx, query,
		book.Title,
		book.AuthorID,
		book.Year,
		book.ISBN,
		book.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *pgDB) DelBookById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        delete from books where id = $1
    `

	_, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

//здесь и далее действия с авторами

func (m *pgDB) GetListAuthors() ([]*models.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        select id, "name", sirname, biography, birthday
		from authors
        order by name
    `

	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []*models.Author

	for rows.Next() {
		var author models.Author
		err := rows.Scan(
			&author.ID,
			&author.Name,
			&author.Sirname,
			&author.Biography,
			&author.Birthday,
		)
		if err != nil {
			return nil, err
		}

		authors = append(authors, &author)
	}

	return authors, nil
}

func (m *pgDB) GetAuthorById(id int) (*models.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select id, "name", sirname, biography, birthday
		from authors
        where id =$1
    `
	row := m.db.QueryRowContext(ctx, query, id)
	var author models.Author
	err := row.Scan(
		&author.ID,
		&author.Name,
		&author.Sirname,
		&author.Biography,
		&author.Birthday,
	)
	if err != nil {
		return nil, err
	}
	return &author, err
}

func (m *pgDB) InsAuthorById(author models.Author) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        insert into authors("name", sirname, biography, birthday)
		values($1, $2, $3, $4) returning id
    `

	t, err := time.Parse("2006-01-02", author.Birthday)
	if err != nil {
		return 0, err
	}

	var newID int

	err = m.db.QueryRowContext(ctx, query,
		author.Name,
		author.Sirname,
		author.Biography,
		t,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (m *pgDB) UpdAuthorById(author models.Author) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        update authors
		set "name" = $1, sirname = $2, biography = $3, birthday = $4
		where id = $5
    `

	fmt.Println(author.ID, author.Birthday)

	_, err := m.db.ExecContext(ctx, query,
		author.Name,
		author.Sirname,
		author.Biography,
		author.Birthday,
		author.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *pgDB) DelAuthorById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        delete from authors where id = $1
    `

	_, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *pgDB) UpdAuthorBook(author models.Author, book models.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	query := `
        update authors
		set "name" = $1, sirname = $2, biography = $3, birthday = $4
		where id = $5
    `

	_, err = m.db.ExecContext(ctx, query,
		author.Name,
		author.Sirname,
		author.Biography,
		author.Birthday,
		author.ID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	query = `
        update books
		set name = $1, authorid = $2, year = $3, isbn =$4
		where id = $5
    `

	_, err = m.db.ExecContext(ctx, query,
		book.Title,
		book.AuthorID,
		book.Year,
		book.ISBN,
		book.ID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
