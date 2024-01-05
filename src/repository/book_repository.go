package repository

import (
	"database/sql"
	"github.com/IndarMuis/go-gin-example.git/src/model/entity"
)

type BookRepository interface {
	FindAll() ([]*entity.Book, error)
	FindById(id uint) (*entity.Book, error)
	FindByName(name string) ([]*entity.Book, error)
	Save(book *entity.Book) (*entity.Book, error)
	Update(book *entity.Book) (*entity.Book, error)
	Delete(book *entity.Book) (*entity.Book, error)
}

type BookRepositoryImpl struct {
	*sql.DB
}

func (bookRepository *BookRepositoryImpl) FindAll() ([]*entity.Book, error) {
	query := "SELECT id, title, author, category, published_year FROM book"
	rows, err := bookRepository.DB.Query(query)
	
	if err != nil {
		return nil, err
	}

	var books []*entity.Book
	for rows.Next() {
		book := entity.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.PublishedYear)
		if err != nil {
			panic(err)
		}
		books = append(books, &book)
	}

	return books, nil
}

func (bookRepository *BookRepositoryImpl) FindById(id uint) (*entity.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (bookRepository *BookRepositoryImpl) FindByName(name string) ([]*entity.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (bookRepository *BookRepositoryImpl) Save(book *entity.Book) (*entity.Book, error) {
	tx, err := bookRepository.DB.Begin()
	if err != nil {
		panic(err)
	}

	//stmt, err := tx.Prepare("INSERT INTO book(title, author, category, publisherYear) VALUES(?, ?, ?, ?, ?)")
	//if err != nil {
	//	panic(err)
	//}

	result, err := tx.Exec("INSERT INTO book(title, author, category, published_year) VALUES(?, ?, ?, ?)", book.Title, book.Author, book.Category, book.PublishedYear)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return &entity.Book{
		ID:            uint(lastInsertId),
		Title:         book.Title,
		Author:        book.Author,
		Category:      book.Category,
		PublishedYear: book.PublishedYear,
	}, nil
}

func (bookRepository *BookRepositoryImpl) Update(book *entity.Book) (*entity.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (bookRepository *BookRepositoryImpl) Delete(book *entity.Book) (*entity.Book, error) {
	//TODO implement me
	panic("implement me")
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &BookRepositoryImpl{DB: db}
}
