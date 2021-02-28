# Polaris	API

API para geração de informações bancárias para boletos desenvolvida em Golang.

Através dela é possível a geração de dados como código barras e linha digitável. A geração dessas informações está disponível para os bancos **Bradesco** e **Santander**.

### Pré-requisitos

* Instalar/configurar o Golang; 
* Ter alguma IDE para desenvolvimento, por exemplo, VSCode;
* Instalar o Git.

### Download e Iniciando

As instruções abaixo possibilitará você fazer o download e configuração do projeto para sua máquina para fins de estudo, desenvolvimento, testes.

Fazer o clone do projeto do github:
> git clone https://github.com/felipesm/polaris-api.git

No VSCode, acessar o menu `Terminal`, `New Terminal` e na aba terminal digitar:
> go run .

Será feito o download do projeto [`polaris-boleto`](https://github.com/felipesm/polaris-boleto) que é uma dependência. Em seguida a aplicação é iniciada na porta 3000.

### API

A API possui dois endpoints, um para geração do código de barras e outro para geração da linha digitável.

Código Bradesco: 237<br>
Código Santander: 033<br>

**Código de Barras**

* **URL**

  /codigobarras
  
* **Método**

  GET
  
*  **Parâmetros URL**

    | Nome  |  Tipo de dado  |
    | :---: | :---: |
    |  codbanco |  string |
    |  agencia |  integer |
    |  carteira |  integer |
    |  valor |  float |
    |  vencimento |  string no formato: yyyy-MM-dd |
    |  numero |  string |
    |  codbeneficiario |  integer |

*  **Resposta Sucesso**

   Code: 200<br>
   Content:
```json
{
  "codigobanco": "237",
  "codigomoeda": "9",
  "digitoverificador": "9",
  "fatorvencimento": "1207",
  "valor": "0000000195",
  "campolivre": "5338370000309067500053040",
  "codigobarras": "23799120700000001955338370000309067500053040"
}
```
*  **Resposta Erro**

   Code: 400<br>
   Content:
```json
{
  "titulo": "Código do Banco Inválido",
  "mensagem": "Ocorreu um erro ao tentar identificar o banco, pois não existe serviço disponível para o código 123! Informe um código de banco válido.",
  "status": 400
}
```

*  **Notas**

   Para o Bradesco, todos os parâmetros acima são obrigatórios.
   
   Para o Santander, o parâmetro *`agência`* não precisa ser enviado na requisição.
   
   Se desejar que o campo *`vencimento`* no Código de Barras seja zerado, informar **`1990-01-01`** na requisição.
   
   Se desejar que o campo *`valor`* no Código de Barras seja zerado, informar **`0`** na requisição.
   
**Linha Digitável**

* **URL**

  /linhadigitavel
  
* **Método**

  GET
  
*  **Parâmetros URL**

    | Nome  |  Tipo de dado  |
    | :---: | :---: |
    |  codigobarras |  string |

*  **Resposta Sucesso**

   Code: 200<br>
   Content:

```json
{
  "linhadigitavel": "23795338357000030906575000530404912070000000195",
  "linhadigitavelformatada": "23795.33835 70000.309065 75000.530404 9 12070000000195"
}
```
    
*  **Resposta Erro**

   Code: 400<br>
   Content:
```json
{
  "titulo": "Código de Barras Inválido",
  "mensagem": "Ocorreu um erro ao tentar gerar a linha digitável, pois o código de barras é inválido! O código 237991207000000019553383700003090675000530409 não tem o tamanho correto.",
  "status": 400
}
```
