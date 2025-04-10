package src

import (
	"bytes"
	"encoding/json"
	"have/entities"
	"log"
	"net/http"
)

func RequestCarsFetchAPI(rent entities.Rent, method string) {
	URL := "http://localhost:8084/cars/" + method 
	jsonBody, _ := json.Marshal(rent)

	req, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Error creando la petición PATCH: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error al ejecutar la petición PATCH: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("La petición PATCH devolvió el estado: %d", resp.StatusCode)
	} else {
		log.Println("Carro actualizado correctamente mediante PATCH")
	}
}