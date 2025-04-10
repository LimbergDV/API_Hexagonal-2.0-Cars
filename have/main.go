package main

import (
	"encoding/json"
	"log"
	"have/entities"
	"have/src"

	"github.com/joho/godotenv"
)

func main() {
  // Cargar las variables de entorno
  godotenv.Load()
  rabbit := src.NewRabbitMQ()
  
  // Tratamiento de un mensaje
  msgs := rabbit.GetMessages()
  var forever chan struct{}

  go func() {
    for d := range msgs {
        var rent entities.Rent
        err := json.Unmarshal(d.Body, &rent)
        if err != nil {
            log.Printf("Error al decodificar el mensaje: %s", err)
            continue
        }
        log.Printf(" [x] Renta recibido: id_customer=%d, id_car=%d, return_date=%s", rent.Id_Customer, rent.Id_Car, rent.Return_date_rent)
        
        if rent.Return_date_rent != "0000-00-00" {
          src.RequestCarsFetchAPI(rent, "rent/")
        } else {
		  src.RequestCarsFetchAPI(rent, "return/")
        } 
    }
}()

  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  <-forever
}

