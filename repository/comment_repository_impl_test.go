package repository

import (
	belajargolangdatabase "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "Repository@test.com",
		Comment: "Test Repository",
	}

	result, err := CommentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())
	comment, err := CommentRepository.FindById(context.Background(), 25)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())
	comments, err := CommentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
