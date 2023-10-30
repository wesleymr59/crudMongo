package repositories

import (
	"api-livro/src/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.TODO()

func FindAll() []models.Livros {
	var livros []models.Livros
	cur, err := DB.Find(ctx, bson.D{})
	if err = cur.All(context.Background(), &livros); err != nil {
		log.Fatal(err)
	}
	fmt.Println(livros)
	return livros
}

func FindOne(id int) []models.Livros {
	var livros []models.Livros
	cur, err := DB.Find(ctx, bson.D{{Key: "id", Value: id}})
	if err = cur.All(context.Background(), &livros); err != nil {
		log.Fatal(err)
	}
	fmt.Println(livros)
	return livros
}

func InsertLivro(livro []interface{}) {
	DB.InsertMany(ctx, livro)
}

func UpdateLivro(livro models.Livros) {
	filter := bson.D{{Key: "id", Value: livro.Id}}
	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "nome", Value: livro.Nome},
			{Key: "autor", Value: livro.Autor},
		},
	}}

	DB.UpdateOne(ctx, filter, update)
}

func DeleteLivro(id int) {
	DB.DeleteOne(ctx, bson.D{{Key: "id", Value: id}})
}
