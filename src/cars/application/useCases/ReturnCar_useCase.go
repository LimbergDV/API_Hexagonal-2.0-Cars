package application

import "segunda-API-w-rabbit/src/cars/domain"

type ReturnCar struct {
	db domain.ICar
}

func NewReturnCar(db domain.ICar) *ReturnCar {
	return &ReturnCar{db: db}
}

func (uc *ReturnCar) Run(id_car int) (uint,error){
	return uc.db.ReturnCar(id_car)
}