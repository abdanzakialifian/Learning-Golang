package test

import (
	"context"
	"fmt"
	"learning-golang-database/entity"
	"learning-golang-database/repository"
	"testing"
)

func TestInsert(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository2@gmail.com",
		Comment: "Repository 2",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())
	ctx := context.Background()
	result, err := commentRepository.FindById(ctx, 55)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())
	ctx := context.Background()
	results, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range results {
		fmt.Println(comment)
	}
}
