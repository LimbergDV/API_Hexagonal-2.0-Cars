package controllers

import (
	"fmt"
	"net/http"
	"segunda-API-w-rabbit/src/cars/application/services"
	application "segunda-API-w-rabbit/src/cars/application/useCases"
	"segunda-API-w-rabbit/src/cars/domain"
	"segunda-API-w-rabbit/src/cars/infrastructure"

	"segunda-API-w-rabbit/src/cars/infrastructure/routes/validators"

	"github.com/gin-gonic/gin"
)

type ReturnCarController struct {
	app *application.ReturnCar
	service *services.NotifyOfReturnEvent
}

func NewReturnCarController() *ReturnCarController {
	mysql := infrastructure.GetMySQL()
	rabbit := infrastructure.GetRabbitMQ()
	app := application.NewReturnCar(mysql)
	service := services.NewNotifyOfReturnEvent(rabbit)
	return &ReturnCarController{app: app, service: service}
}

func (rc_c *ReturnCarController) Run(c *gin.Context) {
	var rent domain.Rent

	if err := c.ShouldBindJSON(&rent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	fmt.Print(rent)
	
	if err := validators.CheckRent(rent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	rowsAffected, _ := rc_c.app.Run(int(rent.Id_Car))

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false, "error": "No se pudo actualizar el carro: No se entontró la referencia o ocurrió algo más",
		})
		return
	}

	// Notificar de devuelto
	//rc_c.service.Run()

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}