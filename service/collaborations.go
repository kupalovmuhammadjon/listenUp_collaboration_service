package service

import (
	pb "collaboration_service/genproto/collaborations"
	pbU "collaboration_service/genproto/user"
	"collaboration_service/storage/postgres"
	"context"
	"database/sql"
)

type Collaborations struct {
	pb.UnimplementedCollaborationsServer
	CollaborationsRepo *postgres.CollaborationRepo
	Client             pbU.UserManagementClient
}

func NewCollaborations(db *sql.DB, client pbU.UserManagementClient) *Collaborations {
	collaborations := postgres.NewCollaborationRepo(db)
	return &Collaborations{
		CollaborationsRepo: collaborations,
		Client:             client}
}

func (c *Collaborations) GetCollaboratorsByPodcastId(ctx context.Context, id *pb.ID) (*pb.Collaborators, error) {
	res := pb.Collaborators{}

	collaboratorsId, err := c.CollaborationsRepo.GetCollaboratorsByPodcastId(id.Id)
	if err != nil {
		return nil, err
	}

	for _, col := range *collaboratorsId {
		userInfo, err := c.Client.GetUserByID(context.Background(), &pbU.ID{Id: col.UserId})
		if err != nil {
			return nil, err
		}
		colaborator := pb.Collaborator{
			Username: userInfo.Username,
			Email:    userInfo.Email,
			Role:     col.Role,
			JoinedAt: col.JoinedAt,
		}
		res.Collaborators = append(res.Collaborators, &colaborator)
	}
	return &res, nil
}
