package main

import (
	cars "segunda-API-w-rabbit/src/cars/infrastructure"
	carsRoutes "segunda-API-w-rabbit/src/cars/infrastructure/routes"
	"segunda-API-w-rabbit/src/core"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main () {
	godotenv.Load()
	cars.GoDependeces()
	
	
	r := gin.Default()

	core.InitCORS(r)

	carsRoutes.Routes(r)
	
	
	r.Run(":8084")
}