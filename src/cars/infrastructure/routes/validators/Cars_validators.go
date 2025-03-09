package validators

import (
	"errors"
	"segunda-API-w-rabbit/src/cars/domain"
)

func CheckCar(car domain.Car) error {

	if car.Id < 0 {
		return errors.New("El ID del carro tiene que ser mayor o igual a 0")
	}

	if car.Brand == "" {
		return errors.New("La marca del carro no puede estar vacía")
	}

	if car.Model == "" {
		return errors.New("El modelo del carro no puede estar vacío")
	}

	if car.Year < 0 { 
		return errors.New("El año del carro no puede ser menor a 1886")
	}

	if car.Type_Car == "" {
		return errors.New("El tipo de carro no puede estar vacío")
	}

	if car.Plate_number == "" {
		return errors.New("El número de placa no puede estar vacío")
	}

	if car.Price_day <= 0 {
		return errors.New("El precio por día debe ser mayor a 0")
	}

	return nil
}