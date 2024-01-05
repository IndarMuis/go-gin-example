package main

import (
	"github.com/IndarMuis/go-gin-example.git/src/common/config"
	"github.com/IndarMuis/go-gin-example.git/src/controller"
	"github.com/IndarMuis/go-gin-example.git/src/repository"
	"github.com/IndarMuis/go-gin-example.git/src/service/impl"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.NewDatabase()

	// init repository
	bookRepository := repository.NewBookRepository(db)

	// init service
	bookService := impl.NewBookService(bookRepository)

	// init controller
	bookController := controller.NewBookController(bookService)

	app := gin.New()
	app.Use(gin.Recovery())

	// init routes
	bookController.Routes(app)

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
