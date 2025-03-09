package application

import "segunda-API-w-rabbit/src/cars/domain"

type CreateCar struct{
	db domain.ICar
}

func NewCreateCar (db domain.ICar) *CreateCar {
	return &CreateCar{db: db}
}

func (cc *CreateCar) Run (car domain.Car) (uint, error) {
	return cc.db.Save(car)
}