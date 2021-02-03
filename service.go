package main

import (
	"encoding/json"
	"net/http"

	"github.com/polaris-boleto/boletos"
)

// GetLinhaDigitavel - retorna código de barras
func GetLinhaDigitavel(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()

	codigoBarras := params.Get("codigobarras")
	codigoBanco := codigoBarras[0:3]

	boleto := boletos.InstanciarBoleto(codigoBanco)

	linha := boleto.GetLinhaDigitavel(codigoBarras)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(linha)

}
