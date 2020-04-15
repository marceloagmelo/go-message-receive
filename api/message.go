package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/marceloagmelo/go-message-receive/utils"

	"github.com/marceloagmelo/go-message-receive/logger"
	"github.com/marceloagmelo/go-message-receive/models"
	"github.com/marceloagmelo/go-message-receive/variaveis"
)

var api = "go-message/api/v1"

//RecuperarMensagem enviar a mensagem
func RecuperarMensagem(id string) (mensagemRetorno models.Mensagem, erro error) {
	endpoint := variaveis.ApiURL + "/" + api + "/mensagem/" + id

	resposta, err := utils.GetRequest(endpoint)
	if err != nil {
		return mensagemRetorno, err
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == http.StatusOK {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o conteudo recebido", err.Error())
			logger.Erro.Println(mensagem)
			return mensagemRetorno, err
		}
		mensagemRetorno = models.Mensagem{}
		err = json.Unmarshal(corpo, &mensagemRetorno)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter o retorno JSON", err.Error())
			logger.Erro.Println(mensagem)
			return mensagemRetorno, err
		}
	}
	return mensagemRetorno, nil
}

//AtualizarMensagem enviar a mensagem
func AtualizarMensagem(novaMensagem models.Mensagem) (mensagemRetorno models.Mensagem, erro error) {
	endpoint := variaveis.ApiURL + "/" + api + "/mensagem/atualizar"

	resposta, err := utils.PutRequest(endpoint, novaMensagem)
	if err != nil {
		return mensagemRetorno, err
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == http.StatusOK {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o conteudo recebido", err.Error())
			logger.Erro.Println(mensagem)
			return mensagemRetorno, err
		}
		mensagemRetorno = models.Mensagem{}
		err = json.Unmarshal(corpo, &mensagemRetorno)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter o retorno JSON", err.Error())
			logger.Erro.Println(mensagem)
			return mensagemRetorno, err
		}
	}
	return mensagemRetorno, nil
}
