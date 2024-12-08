# Prática da Pós "Go Expert": Multithreading

## Resumo

O objetivo é realizar requisições http concorrentes a duas APIs distintas e exibir o resultado apenas da primeira que responder.
Caso o tempo de resposta seja superior a 1 segundo um erro de _timeout_ deve ser exibido.

## Para rodar

```bash
cd cmd/cli
go run main.go
```

## Exemplo de saída

```bash
✦ ❯ go run main.go 
# Service:            BrasilApi
# CEP:                01153000
# Bairro:             Barra Funda
# Rua:                Rua Vitorino Carmilo
# Cidade:             São Paulo
# UF:                 SP
# 
# Time elapsed: 73ms
```
