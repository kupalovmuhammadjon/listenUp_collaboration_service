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

func (c *Collaborations) RespondInvitation(ctx context.Context, req *pb.CreateCollaboration) (*pb.ID, error) {
	resp, err := c.Collaborations.RespondInvitation(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Collaborations) DeleteCollaboratorByPodcastId(ctx context.Context, req *pb.Ids) (*pb.Void, error) {
	resp, err := c.Collaborations.DeleteCollaboratorByPodcastId(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}