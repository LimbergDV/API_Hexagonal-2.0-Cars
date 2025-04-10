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



type RentCarController struct{
	app *application.RentCar
	service *services.NotifyOfRentEvent
}

func NewRentCarController() *RentCarController {
	mysql := infrastructure.GetMySQL()
	rabbit := infrastructure.GetRabbitMQ()
	app := application.NewRentCar(mysql)
	service := services.NewNotifyOfRent(rabbit)
	return &RentCarController{app: app, service: service}
}

func (rc *RentCarController) Run (c *gin.Context){
	var rent domain.Rent

	if err := c.ShouldBindJSON(&rent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false, "error": "datos inválidos" + err.Error(),
		})
		return
	}

	if err := validators.CheckRent(rent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": false, "error": "datos inválidos" + err.Error(),
		})
		return
	}

	fmt.Print(rent)

	rowsAffected, _ := rc.app.Run(int(rent.Id_Car))

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H {
			"status": false, "error": "No se puede actualizar el carro",
		})
		return
	}

	// rc.service.Run(int(rent.Id_Customer), rent.Return_date_rent)
	c.JSON(http.StatusOK, gin.H {
		"status": "está bien la notificación de la renta",
	})
}
