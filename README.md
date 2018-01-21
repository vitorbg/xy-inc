# **_Projeto XY INC_**

## Realizar o download das seguintes bibliotecas utilizadas:

Framework para HTTP

go get github.com/labstack/echo

go get github.com/dgrijalva/jwt-go

Driver do Banco de Dados Sqlite3 (Precisa do gcc instalado, <https://github.com/mattn/go-sqlite3>)

go get github.com/mattn/go-sqlite3

## Realizar o download do código fonte do xy-inc

go get github.com/vitorbg/xy-inc

## Execução do código

Na pasta $GOPATH/src/github.com/vitorbg/xy-inc digitar o seguinte comando:

go run main.go

O nome da base de dados e a porta do servidor são configuráveis pelo arquivo de configurações "xy-inc.json". Por padrão a porta é 1323 e o nome da base de dados "xy-inc.db". Na primeira execução a base sqlite3 será criada, juntamente com o script de criação da tabela POI.

## Rotas

Todas as rotas são públicas. Pode-se abordar o uso de um middleware para segurança das rotas, como por exemplo, o JWT.

--------------------------------------------------------------------------------

GET localhost/poi

--------------------------------------------------------------------------------

GET localhost/poi/proximidade

QueryParam(x) = Coordenada X, QueryParam(y) = Coordenada Y, QueryParam(dmax) = Distância Máxima

--------------------------------------------------------------------------------

POST localhost/poi

Form application/x-www-form-urlencoded

FormValue(nome) = Nome do POI, FormValue(x) = Coordenanda X do POI, FormValue(y) = Coordenada Y do POI

## Exemplo de consumo http utilizando curl

curl -X GET http://localhost:1323/poi

curl -X GET 'http://localhost:1323/poi/proximidade?x=20&y=10&dmax=10'

curl -X POST http://localhost:1323/poi -F nome=Escola -F x=10 -F y=25
