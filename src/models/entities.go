package models

import "gopkg.in/validator.v2"

type MessageHealthy struct {
	Message string `json:"message"`
}

type Livros struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Autor string `json:"autor"`
}

func ValidaDadosLivro(livro []Livros) error {
	if err := validator.Validate(livro); err != nil {
		return err
	}
	return nil
}
