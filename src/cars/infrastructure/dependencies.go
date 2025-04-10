package infrastructure

import "segunda-API-w-rabbit/src/cars/infrastructure/adapters"

var (
	mysql *MySQL
	rabbitmq *adapters.RabbitMQ
)

func GoMySQL() {
	mysql = NewMySQL()
}

func GetMySQL() *MySQL {
	return mysql 
}

func GetRabbitMQ() *adapters.RabbitMQ {
	return rabbitmq
}