package routes

import (
	"net/http"

	"github.com/marceloagmelo/go-message-receive/controllers"
)

//CarregaRotas de acesso
func CarregaRotas() {
	http.HandleFunc("/health", controllers.Health)
}
