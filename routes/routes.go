package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenclaudiano/CriandoAPi2025/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/carros", controllers.ExibeTodosCarros)
	//r.POST("/Carros", controllers.CriaNovoAluno)
	r.Run()
}
