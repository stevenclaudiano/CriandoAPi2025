package models

type Carros struct {
	Id     string `json:"id"`
	Nome   string `json:"nome"`
	Modelo string `json:"modelo"`
	Ano    string `json:"ano"`
	Preco  string `json:"preco"`
}
