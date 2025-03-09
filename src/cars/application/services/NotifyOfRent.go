package services

import "segunda-API-w-rabbit/src/cars/application/repository"



type NotifyOfRentvent struct {
	rmq repository.IRabbit
}

func NewNotifyOfRent(rmq repository.IRabbit) *NotifyOfRentvent {
	return &NotifyOfRentvent{rmq: rmq}
}

func (r *NotifyOfRentvent) Run(id_customer int, return_date string) {
	r.rmq.NotifyOfRent(id_customer, return_date)
}