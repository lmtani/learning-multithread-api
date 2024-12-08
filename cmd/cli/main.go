package main

import (
	"fmt"
	"github.com/lmtani/learning-multithread-api/configs"
	"github.com/lmtani/learning-multithread-api/pkg/cep"
)

func main() {
	config, err := configs.LoadConfig(".") // relative path to main.go
	if err != nil {
		panic(err)
	}

	viacep := cep.NewViaCep(config.ViaCepApiUrl)
	brasilapi := cep.NewBrasilApi(config.BrasilApiUrl)

	c, err := viacep.GetCep("01153000")
	if err != nil {
		panic(err)
	}

	c2, err := brasilapi.GetCep("01153000")
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Printf("%+v\n", c))
	fmt.Println(fmt.Printf("%+v\n", c2))
}
