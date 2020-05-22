package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/marceloagmelo/go-message-receive/api"
	"github.com/marceloagmelo/go-message-receive/model"

	"github.com/marceloagmelo/go-message-receive/logger"
	"github.com/marceloagmelo/go-message-receive/util"
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
	util.CheckErrFatal(err, "Conectando com o rabbitmq")

	return conn, nil
}

//LerMensagensRabbitMQ no rabbitmq
func LerMensagensRabbitMQ(conn *amqp.Connection) {
	// Abrir o canal
	ch, err := conn.Channel()
	util.CheckErrFatal(err, "Falha ao abrir o canal no rabbitmq")
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
	util.CheckErrFatal(err, "Falha ao declarar a fila no rabbitmq")

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
	util.CheckErrFatal(err, "Falha ao ler as mensagens no rabbitmq")

	util.CheckErrFatal(err, "Falha ao ler as mensagens no rabbitmq")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var mensagem model.Mensagem
			err = json.Unmarshal(d.Body, &mensagem)
			if err != nil {
				mensagemErro := fmt.Sprintf("%s: %s", "Erro ao ler a mensagem com o JSON informado", err.Error())
				logger.Erro.Println(mensagemErro)
			}
			mensagemRecuperada, err := api.RecuperarMensagem(strconv.Itoa(mensagem.ID))
			if err == nil {
				mensagemRecuperada.Status = 2
				novaMensagem, err := api.AtualizarMensagem(mensagemRecuperada)
				if err == nil {
					msg := fmt.Sprintf("Processou a mensagem: %v", novaMensagem.ID)
					logger.Info.Println(msg)
				}
			}
		}
	}()

	mensagem := fmt.Sprintf(" [*] Esperando mensagens da fila: %s", fila)
	logger.Info.Println(mensagem)

	http.ListenAndServe(":8080", nil)

	<-forever
}
