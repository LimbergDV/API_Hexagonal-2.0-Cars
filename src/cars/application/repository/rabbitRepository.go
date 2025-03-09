package repository

type IRabbit interface {
	NotifyOfRent(id_customer int, return_date string)
	NotifyOfReturn(id_customer int)
}