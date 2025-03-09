package controllers

import (
	"net/http"
	application "segunda-API-w-rabbit/src/cars/application/useCases"
	"segunda-API-w-rabbit/src/cars/infrastructure"

	"github.com/gin-gonic/gin"
)

type GetAllCarsController struct {
	app *application.GetAllCars
}

func NewGetAllCarsController() *GetAllCarsController {
	mysql := infrastructure.GetMySQL()
	app := application.NewGetAllCars(mysql)
	return &GetAllCarsController{app: app}
}

func (ctrl *GetAllCarsController) Run(c *gin.Context) {
	res := ctrl.app.Run()

	if len(res) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "error": "No se encontró ningún carro"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"Carros": res})
	}
}