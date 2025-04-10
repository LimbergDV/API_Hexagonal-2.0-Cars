package validators

import (
	"errors"
	"segunda-API-w-rabbit/src/cars/domain"
)

func CheckRent(rent domain.Rent) error {
	if rent.Id_Car <= 0 {
		return errors.New("El id tiene que ser mayor a 0")
	}
	if rent.Id_Customer <= 0 {
		return errors.New("El id tiene que ser mayor mayor a 0")
	}
	if rent.Return_date_rent == "" {
		return errors.New("La fecha de regreso del carro esta vacia")
	}
	return nil
}