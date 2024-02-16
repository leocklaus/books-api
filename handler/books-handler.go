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

func GetBookByIdHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Livro com a id: " + id,
	})
}

func AddNewBookHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Colecao de livros",
	})
}

func EditBookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Livro com a id: " + id,
	})
}

func DeleteBookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Deletanto livro com a id: " + id,
	})
}
