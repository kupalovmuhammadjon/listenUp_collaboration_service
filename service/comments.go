package service

import (
	pb "collaboration_service/genproto/comments"
	"collaboration_service/storage/postgres"
	"context"
	"database/sql"
)

type Comments struct {
	pb.UnimplementedCommentsServer
	Comments *postgres.CommentRepo
}

func NewComments(db *sql.DB) *Comments {
	comments := postgres.NewCommentRepo(db)
	return &Comments{Comments: comments}
}

func (c *Comments) CreateCommentByPodcastId(ctx context.Context, comment *pb.CreateComment) (*pb.ID, error) {
	id, err := c.Comments.CreateCommentByPodcastId(comment)
	return id, err
}
