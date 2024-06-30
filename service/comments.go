package service

import (
	pb "collaboration_service/genproto/comments"
	pbu "collaboration_service/genproto/user"
	"collaboration_service/storage/postgres"
	"context"
	"database/sql"
)

type Comments struct {
	pb.UnimplementedCommentsServer
	Comments   *postgres.CommentRepo
	UserClient *pbu.UserManagementClient
}

func NewComments(db *sql.DB, userClient *pbu.UserManagementClient) *Comments {
	comments := postgres.NewCommentRepo(db)
	return &Comments{Comments: comments, UserClient: userClient}
}

func (c *Comments) GetCommentsByPodcastId(ctx context.Context, id *pb.ID) (*pb.AllComments, error) {
	commentInfo, err := c.Comments.GetCommentsByPodcastId(id)
	if err != nil {
		return nil, err
	}
	allComments := pb.AllComments{}
	for i := 0; i < len(*commentInfo); i++ {
		comment := pb.Comment{}
		user, err := (*c.UserClient).GetUserByID(context.Background(), &pbu.ID{Id: (*commentInfo)[i].UserId})
		if err != nil {
			return nil, err
		}
		comment.Username = user.Username
		comment.Content = (*commentInfo)[i].Content
		comment.CreatedAt = (*commentInfo)[i].CreatedAt
		comment.UpdatedAt = (*commentInfo)[i].UpdatedAt
		allComments.Comments = append(allComments.Comments, &comment)
	}
	return &allComments, nil
}
