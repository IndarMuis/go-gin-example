package service

import "github.com/IndarMuis/go-gin-example.git/src/model/dto"

type BookService interface {
	FindAll() ([]dto.BookResponse, error)
	FindById(id uint) (dto.BookResponse, error)
	FindByName(name string) ([]dto.BookResponse, error)
	Save(book dto.BookRequest) (dto.BookResponse, error)
	Update(book dto.BookRequest) (dto.BookResponse, error)
	Delete(id uint) (dto.BookResponse, error)
}
