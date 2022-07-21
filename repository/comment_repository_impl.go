package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"jabutech.com/golang_mysql/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (r *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	// Script sql insert data
	script := "INSERT INTO comments(email, comment) VALUES (?,?)"
	// Insert data
	result, err := r.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	// Get last id
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	// Passing ID into struct id
	comment.ID = int32(id)
	// Return comment
	return comment, err
}
func (r *commentRepositoryImpl) FindByID(ctx context.Context, id int32) (entity.Comment, error) {
	// Use struct comment
	comment := entity.Comment{}
	// Script for select comment by id
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := r.DB.QueryContext(ctx, script, id)
	if err != nil {
		return comment, err
	}
	// Close rows after all process scan done
	defer rows.Close()

	// If rows.Next() is true
	if rows.Next() {
		// Yes,
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// No
		fmtString := fmt.Sprintf("Id %v Not Found", id)
		return comment, errors.New(fmtString)
	}
}
func (r *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	var comments []entity.Comment
	// Script for Find all comments
	script := "SELECT id, email, comment FROM comments"
	rows, err := r.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	// Close rows after all process scan done
	defer rows.Close()

	// Iteration rows.Next()
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
