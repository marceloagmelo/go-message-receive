package main

import (
	"github.com/marceloagmelo/go-message-receive/lib"
	"github.com/marceloagmelo/go-message-receive/logger"
	"github.com/marceloagmelo/go-message-receive/routes"
)

const (
	fila string = "go-rabbitmq"
)

func main() {
	routes.CarregaRotas()
	logger.Info.Println("Listen 8080...")
	conn, _ := lib.ConectarRabbitMQ()
	defer conn.Close()

	lib.LerMensagensRabbitMQ(conn)
}
