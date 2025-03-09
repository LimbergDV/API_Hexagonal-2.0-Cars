package application

import "segunda-API-w-rabbit/src/cars/domain"

type GetAllCars struct {
	db domain.ICar
}

func NewGetAllCars(db domain.ICar) *GetAllCars {
	return &GetAllCars{db: db}
}

func (lc *GetAllCars) Run () []domain.Car {
	return lc.db.GetAll()
}