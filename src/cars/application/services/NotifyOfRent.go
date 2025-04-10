package services

import "segunda-API-w-rabbit/src/cars/application/repository"



type NotifyOfRentEvent struct {
	rmq repository.IRabbit
}

func NewNotifyOfRent(rmq repository.IRabbit) *NotifyOfRentEvent {
	return &NotifyOfRentEvent{rmq: rmq}
}

func (r *NotifyOfRentEvent) Run(id_customer int, return_date string) {
	r.rmq.NotifyOfRent(id_customer, return_date)
}