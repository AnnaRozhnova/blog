package repository

import (
	"fmt"

	"github.com/AnnaRozhnova/blog"
	"github.com/jmoiron/sqlx"
)

type CommentPostgres struct {
	db *sqlx.DB
}

func NewCommentPostgres(db *sqlx.DB) *CommentPostgres {
	return &CommentPostgres{db: db}
}

func (r *CommentPostgres) Create(comment blog.Comment) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (body, username, post_id) values ($1, $2, $3) RETURNING id", commentsTable)
	row := r.db.QueryRow(query, comment.Body, comment.Username, comment.PostId)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r *CommentPostgres) GetByPostId(postId int) ([]blog.Comment, error) {
	var comments []blog.Comment
	query := fmt.Sprintf("SELECT * FROM %s WHERE post_id=$1", commentsTable)
	err := r.db.Select(&comments, query, postId)
	
	fmt.Println("Postgers COMMENTS: ", comments)
	return comments, err
}