package service

import (
	pb "collaboration_service/genproto/collaborations"
	"collaboration_service/storage/postgres"
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

func (c *Collaborations) GetCollaboratorsByPodcastId(id *pb.ID) (*pb.Collaborator, error) {
	collaborators := *pb.Collaborator{}
}
