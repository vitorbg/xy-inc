# ***Projeto XY INC***

## Realizar o download do código fonte do xy-inc

go get github.com/vitorbg/xy-inc

## Realizar o download das seguintes bibliotecas utilizadas:

Framework para HTTP

go get github.com/labstack/echo

Driver do Banco de Dados Sqlite3

go get github.com/mattn/go-sqlite3

## Execução do código

Na pasta $GOPATH/src/github.com/vitorbg/xy-inc digitar o seguinte comando:

go run main.go

O nome da base de dados e a porta do servidor são configuráveis pelo arquivo de configurações
"xy-inc.json". Por padrão a porta é 1323 e o nome da base de dados "xy-inc.db".

## Rotas
Todas as rotas são públicas. Pode-se abordar o uso de um middleware para segurança das rotas, como por exemplo, o JWT.

ROTA localhost/poi


ROTA localhost/poi



ROTA localhost/poi/proximidade
