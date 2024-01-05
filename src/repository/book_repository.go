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
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
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
