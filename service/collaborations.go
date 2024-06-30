package service

import (
	pb "collaboration_service/genproto/collaborations"
	"collaboration_service/storage/postgres"
	"context"
	"database/sql"
)

type Collaborations struct {
	pb.UnimplementedCollaborationsServer
	Collaborations *postgres.CollaborationRepo
}

func NewCollaborations(db *sql.DB) *Collaborations {
	collaborations := postgres.NewCollaborationRepo(db)
	return &Collaborations{Collaborations: collaborations}
}

func (c *Collaborations) CreateInvitation(ctx context.Context, invitation *pb.CreateInvite) (*pb.ID, error) {
	id, err := c.Collaborations.CreateInvitation(invitation)
	if err != nil {
		return nil, err
	}
	return &pb.ID{Id: id}, err
}

func (c *Collaborations) UpdateCollaboratorByPodcastId(ctx context.Context, clb *pb.UpdateCollaborator) (*pb.Void, error) {
	err := c.Collaborations.UpdateCollaboratorByPodcastId(clb)

	return &pb.Void{}, err
}
