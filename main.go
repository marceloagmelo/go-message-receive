package main

import (
	"github.com/marceloagmelo/go-message-receive/lib"
)

const (
	fila string = "go-rabbitmq"
)

func main() {
	conn, _ := lib.ConectarRabbitMQ()
	defer conn.Close()

	lib.LerMensagensRabbitMQ(conn)
}
