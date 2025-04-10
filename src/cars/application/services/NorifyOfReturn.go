package services

import "segunda-API-w-rabbit/src/cars/application/repository"



type NotifyOfReturnEvent struct {
	rmq repository.IRabbit
}

func NewNotifyOfReturnEvent(rmq repository.IRabbit) *NotifyOfReturnEvent {
	return &NotifyOfReturnEvent{rmq: rmq}
}

func (s *NotifyOfReturnEvent) Run() {
	s.rmq.NotifyOfReturn()
}