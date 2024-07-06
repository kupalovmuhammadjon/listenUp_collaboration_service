package service

import (
	"collaboration_service/config"
	pb "collaboration_service/genproto/collaborations"
	pbp "collaboration_service/genproto/podcasts"
	pbu "collaboration_service/genproto/user"
	"collaboration_service/pkg"
	"collaboration_service/storage/postgres"
	"context"
	"database/sql"
	"fmt"
)

type Collaborations struct {
	pb.UnimplementedCollaborationsServer
	Repo          *postgres.CollaborationRepo
	UserClient    pbu.UserManagementClient
	PodcastClient pbp.PodcastsClient
}

func NewCollaborations(db *sql.DB, cfg *config.Config) *Collaborations {
	client := pkg.CreateUserManagementClient(cfg)
	pClient := pkg.NewPodcastsClient(cfg)

	collaborations := postgres.NewCollaborationRepo(db)
	return &Collaborations{
		Repo:          collaborations,
		UserClient:    client,
		PodcastClient: pClient,
	}
}

func (c *Collaborations) CreateInvitation(ctx context.Context, invitation *pb.CreateInvite) (*pb.ID, error) {

	exists, err := c.UserClient.ValidateUserId(ctx, &pbu.ID{Id: invitation.InviteeId})
	if err != nil || !exists.Success {
		return nil, fmt.Errorf("error while validating user id %v", err)
	}
	exists, err = c.UserClient.ValidateUserId(ctx, &pbu.ID{Id: invitation.InviterId})
	if err != nil || !exists.Success {
		return nil, fmt.Errorf("error while validating user id %v", err)
	}
	s, err := c.PodcastClient.ValidatePodcastId(ctx, &pbp.ID{Id: invitation.PodcastId})
	if err != nil || !s.Success {
		return nil, fmt.Errorf("error while validating podcast id %v", err)
	}

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

func (c *Collaborations) GetCollaboratorsByPodcastId(ctx context.Context, id *pb.ID) (*pb.Collaborators, error) {
	res := pb.Collaborators{}

	s, err := c.PodcastClient.ValidatePodcastId(ctx, &pbp.ID{Id: id.Id})
	if err != nil || !s.Success {
		return nil, fmt.Errorf("error while validating podcast id %v", err)
	}

	collaboratorsId, err := c.Repo.GetCollaboratorsByPodcastId(id.Id)
	if err != nil {
		return nil, err
	}

	for _, col := range collaboratorsId {
		userInfo, err := c.UserClient.GetUserByID(ctx, &pbu.ID{Id: col.UserId})
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

func (c *Collaborations) ValidateCollaborationId(ctx context.Context, id *pb.ID) (*pb.Exists, error) {
	exists, err := c.Repo.ValidateCollaborationId(id.Id)

	return exists, err
}

func (c *Collaborations) ValidateInvitationId(ctx context.Context, id *pb.ID) (*pb.Exists, error) {
	exists, err := c.Repo.ValidateInvitationId(id.Id)

	return exists, err
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

func (c *Collaborations) GetAllPodcastsUsersWorkedOn(ctx context.Context, podcastsId *pb.PodcastsId) (*pb.PodcastsId, error) {
	colaborators, err := c.Repo.GetCollaboratorsIdByPodcastsId(&podcastsId.PodcastsId)
	if err != nil {
		return nil, err
	}

	podcastsIdToReturn, err := c.Repo.GetPodcastsIdByCollaboratorsId(colaborators)
	return &pb.PodcastsId{PodcastsId: *podcastsIdToReturn}, err
}

func (c *Collaborations) CreateOwner(ctx context.Context, collab *pb.CreateAsOwner) (*pb.ID, error) {
	colaborators, err := c.Repo.CreateOwner(collab)
	if err != nil {
		return nil, err
	}

	return &pb.ID{Id: colaborators}, err
}
