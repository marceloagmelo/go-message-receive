package lib

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-message-receive/api"

	"github.com/marceloagmelo/go-message-receive/logger"
	"github.com/marceloagmelo/go-message-receive/utils"
	"github.com/streadway/amqp"
)

const (
	fila string = "go-rabbitmq"
)

//ConectarRabbitMQ no rabbitmq
func ConectarRabbitMQ() (*amqp.Connection, error) {
	// Conectar com o rabbitmq
	var connectionString = fmt.Sprintf("amqp://%s:%s@%s:%s%s", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), os.Getenv("RABBITMQ_HOSTNAME"), os.Getenv("RABBITMQ_PORT"), os.Getenv("RABBITMQ_VHOST"))
	conn, err := amqp.Dial(connectionString)
	utils.CheckErrFatal(err, "Conectando com o rabbitmq")

	return conn, nil
}

//LerMensagensRabbitMQ no rabbitmq
func LerMensagensRabbitMQ(conn *amqp.Connection) {
	// Abrir o canal
	ch, err := conn.Channel()
	utils.CheckErrFatal(err, "Falha ao abrir o canal no rabbitmq")
	defer ch.Close()

	// Declarara fila
	q, err := ch.QueueDeclare(
		fila,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	utils.CheckErrFatal(err, "Falha ao declarar a fila no rabbitmq")

	// Ler mensagens
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.CheckErrFatal(err, "Falha ao ler as mensagens no rabbitmq")

	utils.CheckErrFatal(err, "Falha ao ler as mensagens no rabbitmq")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			strID := utils.BytesToString(d.Body)
			mensagem, _ := api.RecuperarMensagem(strID)
			mensagem.Status = 2
			novaMensagem, _ := api.AtualizarMensagem(mensagem)

			msg := fmt.Sprintf("Processou a mensagem: %v", novaMensagem.ID)
			logger.Info.Println(msg)

		}
	}()

	mensagem := fmt.Sprintf(" [*] Esperando mensagens da fila: %s", fila)
	logger.Info.Println(mensagem)

	<-forever
}
