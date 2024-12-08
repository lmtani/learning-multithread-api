package main

import (
	"fmt"
	"time"

	"github.com/lmtani/learning-multithread-api/configs"
	"github.com/lmtani/learning-multithread-api/pkg/cep"
)

func main() {
	config, err := configs.LoadConfig(".") // relative path to main.go
	if err != nil {
		panic(err)
	}

	viacep := cep.NewViaCep(config.ViaCepApiUrl, 1000)
	brasilapi := cep.NewBrasilApi(config.BrasilApiUrl, 1000)

	ch1 := make(chan *cep.Cep)
	ch2 := make(chan *cep.Cep)

	go requestCep(viacep, "01153000", ch1)
	go requestCep(brasilapi, "01153000", ch2)

	select {
	case c := <-ch1:
		fmt.Println(prettyPrint(c, "ViaCep"))
	case c := <-ch2:
		fmt.Println(prettyPrint(c, "BrasilApi"))
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}

func requestCep(service cep.Service, cep string, ch chan<- *cep.Cep) {
	c, err := service.GetCep(cep)
	if err != nil {
		panic(err)
	}
	ch <- c
}

func prettyPrint(c *cep.Cep, service string) string {
	format := "%-19s %s\n"
	return fmt.Sprintf(
		format+format+format+format+format+format+"\nTime elapsed: %dms",
		"Service:", service,
		"CEP:", c.Cep,
		"Bairro:", c.Bairro,
		"Rua:", c.Rua,
		"Cidade:", c.Cidade,
		"UF:", c.Uf,
		c.TimeElapsed,
	)
}
