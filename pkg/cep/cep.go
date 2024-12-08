package cep

// Cep represents a Brazilian postal code
type Cep struct {
	Cep    string `json:"cep"`
	Bairro string `json:"neighborhood"`
	Rua    string `json:"street"`
	Cidade string `json:"city"`
	Uf     string `json:"state"`
}

type Service interface {
	GetCep(cep string) (*Cep, error)
}
