package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenclaudiano/CriandoAPi2025/database"
	"github.com/stevenclaudiano/CriandoAPi2025/models"
)

func ExibeTodosCarros(c *gin.Context) {
	var carro []models.Carros
	database.DB.Find(&carro)
	c.JSON(200, carro)
}

/*func CriaNovoAluno(c *gin.Context) {
	var Carros models.Carros
	if err := c.ShouldBindJSON(&Carros); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&Carros)
	c.JSON(http.StatusOK, Carros)
}*/
