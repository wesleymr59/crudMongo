package services

import "api-livro/src/models"

func MountPayloadinsertMongo(livros []models.Livros) []interface{} {
	newValue := make([]interface{}, len(livros))

	for i := range livros {
		newValue[i] = livros[i]
	}

	return newValue
}
