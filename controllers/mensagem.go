package controllers

import (
	"fmt"
	"net/http"

	"github.com/marceloagmelo/go-message-receive/lib"
	"github.com/marceloagmelo/go-message-receive/variaveis"
)

type retorno struct {
	Status string `json:"mensagem"`
}

//Health testa conexão com o mysql e rabbitmq
func Health(w http.ResponseWriter, r *http.Request) {
	dataHoraFormatada := variaveis.DataHoraAtual.Format(variaveis.DataFormat)

	conn, err := lib.ConectarRabbitMQ()
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao conectar com o rabbitmq", err)
		respondError(w, http.StatusInternalServerError, mensagem)
		return
	}
	defer conn.Close()

	retorno := retorno{}
	retorno.Status = fmt.Sprintf("OK [%v] !", dataHoraFormatada)

	respondJSON(w, http.StatusOK, retorno)
}
