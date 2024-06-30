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

func (c *Collaborations) CreateInvitation(ctx context.Context, in *pb.CreateInvite) (*pb.ID, error) {

}
