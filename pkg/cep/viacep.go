package cep

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type ViaCep struct {
	url     string
	timeout time.Duration
}

func NewViaCep(host string, timeout int) *ViaCep {
	return &ViaCep{url: host, timeout: time.Duration(timeout) * time.Millisecond}
}

type ViaCepOutput struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (v *ViaCep) GetCep(cep string) (*Cep, error) {
	start := time.Now()
	route := fmt.Sprintf("%s/ws/%s/json/", v.url, cep)
	ctx, cancel := context.WithTimeout(context.Background(), v.timeout)
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

	var p *ViaCepOutput
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, err
	}

	c := &Cep{
		Cep:         p.Cep,
		Bairro:      p.Bairro,
		Rua:         p.Logradouro,
		Cidade:      p.Localidade,
		Uf:          p.Uf,
		TimeElapsed: time.Since(start).Milliseconds(),
	}

	return c, nil
}
