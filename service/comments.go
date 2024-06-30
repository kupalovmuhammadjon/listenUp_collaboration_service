package service

import (
	pb "collaboration_service/genproto/comments"
	"collaboration_service/storage/postgres"
	"context"
	"database/sql"
)

type Comments struct {
	pb.UnimplementedCommentsServer
	Repo *postgres.CommentRepo
}

func NewComments(db *sql.DB) *Comments {
	comments := postgres.NewCommentRepo(db)
	return &Comments{Repo: comments}
}

func (c *Comments) CreateCommentByPodcastId(ctx context.Context, comment *pb.CreateComment) (*pb.ID, error) {
	id, err := c.Repo.CreateCommentByPodcastId(comment)
	return id, err
}
