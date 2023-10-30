package controllers

import (
	"api-livro/src/models"
	"api-livro/src/repositories"
	"api-livro/src/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Healthy_check(c *gin.Context) {
	var message models.MessageHealthy
	message.Message = "Api is runing"
	c.JSON(200, message)
}

func TesteMongo(c *gin.Context) {
	result := repositories.FindAll()
	c.JSON(200, result)
}

func InsertLivro(c *gin.Context) {
	var livro []models.Livros
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := models.ValidaDadosLivro(livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	payloadMongo := services.MountPayloadinsertMongo(livro)
	repositories.InsertLivro(payloadMongo)
	c.JSON(201, "")
}

func UpdateLivro(c *gin.Context) {
	var livro models.Livros
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	repositories.UpdateLivro(livro)
	c.JSON(200, livro)
}

func DeleteLivro(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	fmt.Println(id)
	result := repositories.FindOne(id)
	repositories.DeleteLivro(id)
	c.JSON(200, result)
}
