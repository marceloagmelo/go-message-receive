package util

import (
	"fmt"
	"log"
	"unsafe"
)

// CheckErrFatal checar o erro
func CheckErrFatal(err error, msg string) {
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", msg, err)
		log.Fatalln(mensagem)
	}
}

//BytesToString converter bytes para string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//InBetween Intervalo de números
func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}
