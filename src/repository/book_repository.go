package repository

import (
	"fmt"
	"github.com/IndarMuis/go-gin-example.git/src/model/entity"
	"gorm.io/gorm"
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
	*gorm.DB
}

func (bookRepository *BookRepositoryImpl) FindAll() ([]*entity.Book, error) {
	var books []*entity.Book

	//query := "SELECT id, title, author_id, category, published_year FROM book"
	rows := bookRepository.DB.Debug().Select("id, title, author_id, category_id, published_year").Find(&books)
	if rows.Error != nil {
		return nil, rows.Error
	}

	//var books []*entity.Book
	//for rows.Next() {
	//	book := entity.Book{}
	//	err = rows.Scan(&book.ID, &book.Title, &book.AuthorId, &book.Category, &book.PublishedYear)
	//	if err != nil {
	//		panic(err)
	//	}
	//	books = append(books, &book)
	//}

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
	tx := bookRepository.DB.Debug().Begin()

	// insert into book table
	saveBook := tx.Exec("INSERT INTO book (title, author_id, category_id, published_year) VALUES (?, ?, ?, ?)",
		book.Title, book.AuthorId, book.CategoryId, book.PublishedYear)

	if saveBook.Error != nil {
		tx.Rollback()
		return nil, saveBook.Error
	}

	var lastInsertId uint
	saveBook.Select("id").Last(&entity.Book{}).Scan(&lastInsertId)
	fmt.Println(lastInsertId)

	// insert into book_category table
	saveBookCategory := tx.Exec("INSERT INTO book_category (book_id, category_id) VALUES (?, ?)",
		lastInsertId, book.CategoryId)

	if saveBookCategory.Error != nil {
		tx.Rollback()
		return nil, saveBookCategory.Error
	}

	tx.Commit()
	//tx.Rollback()

	return &entity.Book{
		ID:            lastInsertId,
		Title:         book.Title,
		AuthorId:      book.AuthorId,
		CategoryId:    book.CategoryId,
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

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{DB: db}
}
