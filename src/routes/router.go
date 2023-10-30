package routes

import (
	"api-livro/src/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/", controllers.Healthy_check)
	r.GET("/testezin", controllers.TesteMongo)
	r.POST("/livros", controllers.InsertLivro)
	r.PUT("/livros/:id", controllers.UpdateLivro)
	r.DELETE("/livros/:id", controllers.DeleteLivro)
	r.Run()

}
