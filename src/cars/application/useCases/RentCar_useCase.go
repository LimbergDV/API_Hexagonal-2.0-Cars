package application

import "segunda-API-w-rabbit/src/cars/domain"

type RentCar struct {
	db domain.ICar
}

func NewRentCar(db domain.ICar) *RentCar {
	return &RentCar{db: db}
}

func (uc *RentCar) Run(id_car int) (uint, error) {
	return uc.db.RentCar(id_car)
}