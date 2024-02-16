package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leocklaus/books-api/handler"
)

func InitializeRoutes(r *gin.Engine) {

	v1 := r.Group("/api/v1")

	v1.GET("/books", handler.GetBookCollectionHandler)

	v1.GET("/books/:id", handler.GetBookByIdHandler)

	v1.POST("/books", handler.AddNewBookHandler)

	v1.PUT("/books/:id", handler.EditBookHandler)

	v1.DELETE("/books/:id", handler.DeleteBookHandler)

}
