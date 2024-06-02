package impl

import (
	"github.com/IndarMuis/go-gin-example.git/src/model/dto"
	"github.com/IndarMuis/go-gin-example.git/src/model/entity"
	"github.com/IndarMuis/go-gin-example.git/src/repository"
	"github.com/IndarMuis/go-gin-example.git/src/service"
)

type BookServiceImpl struct {
	bookRepository repository.BookRepository
}

func (service *BookServiceImpl) FindAll() ([]dto.BookResponse, error) {
	resultBooks, err := service.bookRepository.FindAll()
	if err != nil {
		panic(err)
	}

	var books []dto.BookResponse
	for _, book := range resultBooks {
		books = append(books, dto.BookResponse{
			ID:            book.ID,
			Title:         book.Title,
			AuthorId:      book.AuthorId,
			Category:      book.Category,
			CategoryId:    book.CategoryId,
			PublishedYear: book.PublishedYear,
		})
	}

	return books, nil
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
	newBook := entity.Book{
		Title:         book.Title,
		AuthorId:      book.AuthorId,
		CategoryId:    book.CategoryId,
		PublishedYear: book.PublishedYear,
	}

	result, err := service.bookRepository.Save(&newBook)
	if err != nil || result == nil {
		panic(err)
	}

	return dto.BookResponse{
		ID:            result.ID,
		Title:         result.Title,
		Author:        result.Author,
		AuthorId:      result.AuthorId,
		Category:      result.Category,
		CategoryId:    result.CategoryId,
		PublishedYear: result.PublishedYear,
	}, nil
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
