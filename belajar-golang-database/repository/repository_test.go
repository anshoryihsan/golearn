package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "test@gmail.com",
		Comment: "test repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comment {
		fmt.Println(comment)
	}
}
