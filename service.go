package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/felipesm/polaris-boleto/boletos"
	errorutil "github.com/felipesm/polaris-boleto/erro"
)

// getCodigoBarras - retorna código de barras
func getCodigoBarras(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()

	codigoBarras, erro := gerarCodigoBarras(params)

	w.Header().Set("Content-Type", "application/json")

	if erro.Status != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(erro)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(codigoBarras)
	}
}

func gerarCodigoBarras(params url.Values) (boletos.CodigoBarras, errorutil.Erro) {

	var codigoBarras boletos.CodigoBarras
	var erro errorutil.Erro

	codigoBanco := params.Get("codbanco")
	agencia, err := strconv.Atoi(params.Get("agencia"))
	carteira, err := strconv.Atoi(params.Get("carteira"))
	valor, err := strconv.ParseFloat(params.Get("valor"), 64)
	vencimento := params.Get("vencimento")
	nossoNumero := params.Get("numero")
	codigobeneficiario, err := strconv.Atoi(params.Get("codbeneficiario"))

	boleto, erro := boletos.InstanciarBoleto(codigoBanco)

	if erro.Status != 0 {
		return codigoBarras, erro
	}

	if err != nil {
		return codigoBarras, errorutil.Erro{Titulo: "Dados Inválidos", Mensagem: "Um ou mais dados informados são inválidos. Favor verificar as informações e solicitar novamente.", Status: 400}
	}

	boleto.SetCodigo()
	boleto.SetAgencia(int32(agencia))
	boleto.SetCarteira(int16(carteira))
	boleto.SetFatorVencimento(vencimento, vencimento == "1990-01-01")
	boleto.SetValorBoleto(valor, valor == 0)
	boleto.SetNossoNumero(nossoNumero)
	boleto.SetCodigoBeneficiario(int32(codigobeneficiario))

	return boleto.GetCodigoBarras()
}

// getLinhaDigitavel - retornar linha digitavel
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

// gerarLinhaDigitavel - realiza a geração da linha digitavel através da biblioteca polaris
func gerarLinhaDigitavel(params url.Values) (boletos.LinhaDigitavel, errorutil.Erro) {

	var linha boletos.LinhaDigitavel
	var erro errorutil.Erro

	codigoBarras := params.Get("codigobarras")
	codigoBanco := codigoBarras[0:3]

	boleto, erro := boletos.InstanciarBoleto(codigoBanco)

	if erro.Status != 0 {
		return linha, erro
	}

	return boleto.GetLinhaDigitavel(codigoBarras)
}
