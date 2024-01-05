package controller

import (
	"github.com/IndarMuis/go-gin-example.git/src/model"
	"github.com/IndarMuis/go-gin-example.git/src/model/dto"
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

	route.GET("/books", controller.FindAll)
	route.POST("/books", controller.Save)
	route.GET("/books/:id", controller.FindById)
	//route.GET("/books/:findByName", controller.FindByName)
	route.PUT("/books", controller.Update)
	route.DELETE("/books", controller.Delete)
}

func (controller *BookController) FindAll(context *gin.Context) {
	response, err := controller.BookService.FindAll()
	if err != nil {
		context.JSON(500, model.ResponseTemplate{
			Code:    500,
			Message: "INTERNAL_SERVER_ERROR",
		})
	}

	context.JSON(200, model.ResponseTemplate{
		Code:    200,
		Message: "OK",
		Data:    response,
	})
}

func (controller *BookController) Save(context *gin.Context) {
	var bookRequest dto.BookRequest
	err := context.ShouldBindJSON(&bookRequest)
	if err != nil {
		context.JSON(400, model.ResponseTemplate{
			Code:    400,
			Message: "BAD_REQUEST",
		})
	}

	bookResponse, err := controller.BookService.Save(bookRequest)

	if err != nil {
		context.JSON(500, model.ResponseTemplate{
			Code:    500,
			Message: "INTERNAL_SERVER_ERROR",
		})
	}

	context.JSON(201, model.ResponseTemplate{
		Code:    201,
		Message: "CREATED",
		Data:    bookResponse,
	})
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
