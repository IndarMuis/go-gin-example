package controller

import (
	"github.com/IndarMuis/go-gin-example.git/src/service"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return BookController{BookService: bookService}
}

func (controller *BookController) Routes(app *gin.Engine) {
	route := app.Group("/api/v1")
	route.GET("/books", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"data": "Book controller",
		})
	})

	route.POST("/books", controller.Save)
	route.GET("/books:id", controller.FindById)
	//route.GET("/books/findByName", controller.FindByName)
	route.PUT("/books", controller.Update)
	route.DELETE("/books", controller.Delete)
}

func (controller *BookController) FindAll(context *gin.Context) {
	panic("implement me")
}

func (controller *BookController) Save(context *gin.Context) {
	panic("implement me")
}

func (controller *BookController) FindById(context *gin.Context) {
	panic("implement me")
}

func (controller *BookController) FindByName(context *gin.Context) {
	panic("implement me")
}

func (controller *BookController) Update(context *gin.Context) {
	panic("implement me")
}

func (controller *BookController) Delete(context *gin.Context) {
	panic("implement me")
}
