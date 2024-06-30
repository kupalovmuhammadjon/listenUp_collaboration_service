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
	Repo   *postgres.CollaborationRepo
	Client pbU.UserManagementClient
}

func NewCollaborations(db *sql.DB, client pbU.UserManagementClient) *Collaborations {
	collaborations := postgres.NewCollaborationRepo(db)
	return &Collaborations{
		Repo:   collaborations,
		Client: client}
}

func (c *Collaborations) GetCollaboratorsByPodcastId(ctx context.Context, id *pb.ID) (*pb.Collaborators, error) {
	res := pb.Collaborators{}

	collaboratorsId, err := c.Repo.GetCollaboratorsByPodcastId(id.Id)
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

func (c *Collaborations) UpdateCollaboratorByPodcastId(ctx context.Context, clb *pb.UpdateCollaborator) (*pb.Void, error) {
	err := c.Repo.UpdateCollaboratorByPodcastId(clb)

	return &pb.Void{}, err
}

func (c *Collaborations) DeleteCollaboratorByPodcastId(ctx context.Context, req *pb.Ids) (*pb.Void, error) {
	resp, err := c.Repo.DeleteCollaboratorByPodcastId(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Collaborations) CreateInvitation(ctx context.Context, invitation *pb.CreateInvite) (*pb.ID, error) {
	id, err := c.Repo.CreateInvitation(invitation)
	if err != nil {
		return nil, err
	}
	return &pb.ID{Id: id}, err
}

func (c *Collaborations) RespondInvitation(ctx context.Context, req *pb.CreateCollaboration) (*pb.ID, error) {
	resp, err := c.Repo.RespondInvitation(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
