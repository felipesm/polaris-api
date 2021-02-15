package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/linhadigitavel", getLinhaDigitavel)
	http.HandleFunc("/codigobarras", getCodigoBarras)
	porta := ":3000"
	log.Println("Polaris API iniciada na porta", porta)
	log.Fatal(http.ListenAndServe(porta, nil))
}
