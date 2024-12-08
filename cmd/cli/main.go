package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lmtani/learning-multithread-api/configs"
	"github.com/lmtani/learning-multithread-api/internal/dto"
	"net/http"
	"time"
)

func main() {
	config, err := configs.LoadConfig(".") // relative path to main.go
	if err != nil {
		panic(err)
	}

	cep, err := GetViaCep(config.ViaCepApiUrl, "01153000")
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Printf("%+v\n", cep))

	cep2, err := GetBrasilApi(config.BrasilApiUrl, "01153000")
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Printf("%+v\n", cep2))
}

func GetViaCep(h, cepNum string) (*dto.ViaCepOutput, error) {
	route := fmt.Sprintf("%s/ws/%s/json/", h, cepNum)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, route, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request failed")
	}

	var cep *dto.ViaCepOutput
	if err := json.NewDecoder(resp.Body).Decode(&cep); err != nil {
		return nil, err
	}

	return cep, nil
}

func GetBrasilApi(h, cepNum string) (*dto.BrasilApiOutput, error) {
	route := fmt.Sprintf("%s/api/cep/v1/%s", h, cepNum)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, route, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request failed")
	}

	var cep *dto.BrasilApiOutput
	if err := json.NewDecoder(resp.Body).Decode(&cep); err != nil {
		return nil, err
	}

	return cep, nil
}
