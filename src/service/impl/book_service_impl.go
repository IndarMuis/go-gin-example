package impl

import (
	"github.com/IndarMuis/go-gin-example.git/src/model/dto"
	"github.com/IndarMuis/go-gin-example.git/src/repository"
	"github.com/IndarMuis/go-gin-example.git/src/service"
)

type BookServiceImpl struct {
	bookRepository repository.BookRepository
}

func (service *BookServiceImpl) FindAll() ([]dto.BookResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (service *BookServiceImpl) FindById(id uint) (dto.BookResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (service *BookServiceImpl) FindByName(name string) ([]dto.BookResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (service *BookServiceImpl) Save(book dto.BookRequest) (dto.BookResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (service *BookServiceImpl) Update(book dto.BookRequest) (dto.BookResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (service *BookServiceImpl) Delete(id uint) (dto.BookResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewBookService(bookRepository repository.BookRepository) service.BookService {
	return &BookServiceImpl{bookRepository: bookRepository}
}
