package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookCollectionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Colecao de livros",
	})
}

func GetBookByIdHandler(c *gin.Context, id string) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Livro com a id: " + id,
	})
}

func AddNewBookHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Colecao de livros",
	})
}
