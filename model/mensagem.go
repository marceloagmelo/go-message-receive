package model

//Mensagem estrutura de mensagem
type Mensagem struct {
	ID     int    `db:"id" json:"id"`
	Titulo string `db:"titulo" json:"titulo"`
	Texto  string `db:"texto" json:"texto"`
	Status int    `db:"status" json:"status"`
}

//MensagemHealth retorno do health
type MensagemHealth struct {
	Mensagem string `json:"mensagem"`
}
