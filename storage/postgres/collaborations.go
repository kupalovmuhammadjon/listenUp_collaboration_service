package postgres

import (
	pb "collaboration_service/genproto/collaborations"
	"database/sql"
	"github.com/google/uuid"
)

type CollaborationRepo struct {
	Db *sql.DB
}

func NewCollaborationRepo(db *sql.DB) *CollaborationRepo {
	return &CollaborationRepo{Db: db}
}

func (c *CollaborationRepo) CreateInvitation(invitation *pb.CreateInvite) (string, error) {
	query := `
	insert into 
	    invitations (id, podcast_id, inviter_id, invitee_id)
	values ($1, $2, $3, $4)
`
	tx, err := c.Db.Begin()
	if err != nil {
		return "", err
	}
	id := uuid.NewString()
	_, err = tx.Exec(query, id, invitation.PodcastId, invitation.InviterId, invitation.InviteeId)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *CollaborationRepo) UpdateCollaboratorByPodcastId(clb pb.UpdateCollaborator) error {

	query := `
	update 
	    set
	     	podcast_id = $1,
	    	user_id = $2, 
	    	role = $3
	from
	    collaborations
	where
	    id = $4
`
	tx, err := c.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, clb.PodcastId, clb.UserId, clb.Role, clb.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
