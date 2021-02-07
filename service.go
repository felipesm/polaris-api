package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/polaris-boleto/boletos"
)

// GetCodigoBarras - retorna código de barras
func GetCodigoBarras(w http.ResponseWriter, r *http.Request) {

	// params := r.URL.Query()

	// codigoBanco := params.Get("codbanco")
	// valor := params.Get("valor")
	// vencimento := params.Get("vencimento")
	// carteira := params.Get("carteira")
	// nossoNumero := params.Get("numero")
	// codigobeneficiario := params.Get("codbeneficiario")

	// boleto := boletos.InstanciarBoleto(codigoBanco)

}

// gerarLinhaDigitavel - realiza a geração da linha digitavel através da biblioteca polaris
func gerarLinhaDigitavel(params url.Values) (boletos.LinhaDigitavel, boletos.Erro) {

	var linha boletos.LinhaDigitavel
	var erro boletos.Erro

	codigoBarras := params.Get("codigobarras")
	codigoBanco := codigoBarras[0:3]

	boleto, erro := boletos.InstanciarBoleto(codigoBanco)

	if erro.Status != 0 {
		return linha, erro
	}

	return boleto.GetLinhaDigitavel(codigoBarras)
}

// GetLinhaDigitavel - retornar linha digitavel
func getLinhaDigitavel(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()

	linha, erro := gerarLinhaDigitavel(params)

	w.Header().Set("Content-Type", "application/json")

	if erro.Status != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(erro)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(linha)
	}
}
