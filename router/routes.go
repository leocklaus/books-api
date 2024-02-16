package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {

	v1 := r.Group("/api/v1")

	v1.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Colecao de livros",
		})
	})

	v1.GET("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Livro com a id: " + id,
		})
	})

	v1.POST("/books", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "poing",
		})
	})

	v1.PUT("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Livro com o id: " + id,
		})
	})

	v1.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Deletando livro com id: " + id,
		})
	})

}
