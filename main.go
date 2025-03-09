package main

import (
	cars "segunda-API-w-rabbit/src/cars/infrastructure"
	carsRoutes "segunda-API-w-rabbit/src/cars/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func main () {
	cars.GoMySQL()
	

	r := gin.Default()

	carsRoutes.Routes(r)
	
	
	r.Run(":8080")
}