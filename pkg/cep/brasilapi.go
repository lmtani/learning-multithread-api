package cep

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type BrasilApiOutput struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type BrasilApi struct {
	url     string
	timeout time.Duration
}

func NewBrasilApi(host string, timeout int) *BrasilApi {
	return &BrasilApi{url: host, timeout: time.Duration(timeout) * time.Millisecond}
}

func (b *BrasilApi) GetCep(cep string) (*Cep, error) {
	start := time.Now()
	route := fmt.Sprintf("%s/api/cep/v1/%s", b.url, cep)
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

	var p *BrasilApiOutput
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, err
	}

	c := &Cep{
		Cep:         p.Cep,
		Bairro:      p.Neighborhood,
		Rua:         p.Street,
		Cidade:      p.City,
		Uf:          p.State,
		TimeElapsed: time.Since(start).Milliseconds(),
	}
	return c, nil
}
