package domain

import "fmt"

type Car struct {
		Id int32
		Brand string
		Model string
		Year int
		Type_Car string
		Plate_number string
		Price_day int
		Available bool
}

func NewCar(Brand string, Model string, Year int, Type_Car string, Plate_number string, Price_day int, Available bool) *Car {
	return &Car{Brand: Brand, Model: Model, Year: Year, Type_Car: Type_Car, Plate_number: Plate_number, Price_day: Price_day, Available: Available}
}

func (c *Car) Show() string {
	return fmt.Sprintf("{id: %d, Brand: %s, Model: %s, Year: %s, Type_Car: %d, Plate_number: %d, Price_day: %d, Available: %s}",
		c.Id, c.Brand, c.Model, c.Year, c.Type_Car, c.Plate_number, c.Price_day, c.Available)
}