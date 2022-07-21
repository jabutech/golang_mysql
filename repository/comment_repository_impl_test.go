package repository

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"jabutech.com/golang_mysql/config"
	"jabutech.com/golang_mysql/entity"
)

func TestInsertComment(t *testing.T) {
	// Open connection
	db := config.GetConnection()

	// Call repository
	commentRepository := NewCommentRepository(db)

	// Create context
	ctx := context.Background()
	// Sample data
	comment := entity.Comment{
		Email:   "test@repository.com",
		Comment: "Test Repository",
	}

	newComment, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, newComment.ID)
	assert.Equal(t, comment.Email, newComment.Email)
	assert.Equal(t, comment.Comment, newComment.Comment)
}

func TestFindBy(t *testing.T) {
	// Open connection
	db := config.GetConnection()

	// Use repository
	commentRepository := NewCommentRepository(db)
	// Create context
	ctx := context.Background()

	// Find by id
	comment, err := commentRepository.FindByID(ctx, 1)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, int32(1), comment.ID)
	assert.Equal(t, "test@repository.com", comment.Email)
	assert.Equal(t, "Test Repository", comment.Comment)
}

func TestFindAll(t *testing.T) {
	// Open connection
	db := config.GetConnection()

	// Use repository
	commentRepository := NewCommentRepository(db)
	// Create context
	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, commment := range comments {
		assert.NotNil(t, commment.ID)
		assert.NotNil(t, commment.Email)
		assert.NotNil(t, commment.Comment)
	}
}
