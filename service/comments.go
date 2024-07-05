package service

import (
	"collaboration_service/config"
	pb "collaboration_service/genproto/comments"
	pbp "collaboration_service/genproto/podcasts"
	pbe "collaboration_service/genproto/episodes"
	pbu "collaboration_service/genproto/user"
	"collaboration_service/pkg"
	"collaboration_service/storage/postgres"
	"context"
	"database/sql"
	"fmt"
)

type Comments struct {
	pb.UnimplementedCommentsServer
	Repo          *postgres.CommentRepo
	UserClient    pbu.UserManagementClient
	PodcastClient pbp.PodcastsClient
	EpisodeClient pbe.EpisodesServiceClient
}

func NewComments(db *sql.DB, cfg *config.Config) *Comments {
	client := pkg.CreateUserManagementClient(cfg)
	pClient := pkg.NewPodcastsClient(cfg)
	eClient := pkg.NewEpisodesClient(cfg)
	comments := postgres.NewCommentRepo(db)
	return &Comments{
		Repo:          comments,
		UserClient:    client,
		PodcastClient: pClient,
		EpisodeClient: eClient,
	}
}

func (c *Comments) CreateCommentByPodcastId(ctx context.Context, comment *pb.CreateComment) (*pb.ID, error) {
	exists, err := c.UserClient.ValidateUserId(ctx, &pbu.ID{Id: comment.UserId})
	if err != nil || !exists.Success {
		return nil, fmt.Errorf("error while validating user id %v", err)
	}
	s, err := c.PodcastClient.ValidatePodcastId(ctx, &pbp.ID{Id: comment.PodcastId})
	if err != nil || !s.Success {
		return nil, fmt.Errorf("error while validating podcast id %v", err)
	}

	id, err := c.Repo.CreateCommentByPodcastId(comment)
	return id, err
}

func (c *Comments) CreateCommentByEpisodeId(ctx context.Context, epComment *pb.EpisodeComment) (*pb.ID, error) {
	exists, err := c.UserClient.ValidateUserId(ctx, &pbu.ID{Id: epComment.UserId})
	if err != nil || !exists.Success {
		return nil, fmt.Errorf("error while validating user id %v", err)
	}
	s, err := c.EpisodeClient.ValidateEpisodeId(ctx, &pbe.ID{Id: epComment.EpisodeId})
	if err != nil || !s.Success {
		return nil, fmt.Errorf("error while validating episode id %v", err)
	}

	id, err := c.Repo.CreateEpisodeComment(epComment)

	return &pb.ID{Id: id}, err
}

func (c *Comments) ValidateCommentId(ctx context.Context, id *pb.ID) (*pb.Exists, error) {
	exists, err := c.Repo.ValidateCommentId(id.Id)

	return exists, err
}

func (c *Comments) GetCommentsByPodcastId(ctx context.Context, filter *pb.CommentFilter) (*pb.AllComments, error) {
	commentInfo, err := c.Repo.GetCommentsByPodcastId(filter)
	if err != nil {
		return nil, err
	}
	allComments := pb.AllComments{}
	for i := 0; i < len(commentInfo); i++ {
		comment := pb.Comment{}
		user, err := c.UserClient.GetUserByID(context.Background(), &pbu.ID{Id: (commentInfo)[i].UserId})
		if err != nil {
			return nil, err
		}
		comment.Username = user.Username
		comment.Content = (commentInfo)[i].Content
		comment.CreatedAt = (commentInfo)[i].CreatedAt
		comment.UpdatedAt = (commentInfo)[i].UpdatedAt
		allComments.Comments = append(allComments.Comments, &comment)
	}
	return &allComments, nil
}

func (c *Comments) GetCommentsByEpisodeId(ctx context.Context, filter *pb.CommentFilter) (*pb.AllComments, error) {
	commentInfo, err := c.Repo.GetCommentsByEpisodeId(filter)
	if err != nil {
		return nil, err
	}
	allComments := pb.AllComments{}
	for i := 0; i < len(commentInfo); i++ {
		comment := pb.Comment{}
		user, err := c.UserClient.GetUserByID(context.Background(), &pbu.ID{Id: (commentInfo)[i].UserId})
		if err != nil {
			return nil, err
		}
		comment.Username = user.Username
		comment.Content = (commentInfo)[i].Content
		comment.CreatedAt = (commentInfo)[i].CreatedAt
		comment.UpdatedAt = (commentInfo)[i].UpdatedAt
		allComments.Comments = append(allComments.Comments, &comment)
	}
	return &allComments, nil
}

func (c *Comments) CountComments(ctx context.Context, filter *pb.CountFilter) (*pb.CommentCount, error) {
	count, err := c.Repo.CountComments(filter)

	return count, err
}
