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

func (c *CollaborationRepo) GetCollaboratorsByPodcastId(PodcastId string) (*[]pb.CollaboratorToGet, error) {
	query := `select user_id, role, created_at from collaborations
	where podcast_id = $1`

	collabrators := []pb.CollaboratorToGet{}
	rows, err := c.Db.Query(query, PodcastId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c pb.CollaboratorToGet
		err := rows.Scan(&c.UserId, &c.Role, &c.JoinedAt)
		if err != nil {
			return nil, err
		}
		collabrators = append(collabrators, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &collabrators, nil
}

func (c *CollaborationRepo) UpdateCollaboratorByPodcastId(clb *pb.UpdateCollaborator) error {

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

func (c *CollaborationRepo) DeleteCollaboratorByPodcastId(ids *pb.Ids) (*pb.Void, error) {
	query := `delete from collaborations
	where podcast_id = $1 and user_id = $2`

	tr, err := c.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	_, err = tr.Exec(query, ids.PodcastId, ids.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
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

func (c *CollaborationRepo) RespondInvitation(collab *pb.CreateCollaboration) (*pb.ID, error) {
	query := `
	update invitations
	set status = $1
	where podcast_id = $2 and inviter_id = $3 and invitee_id = $4 and deleted_at = null
	returning id`
	params := []interface{}{collab.Status, collab.Invitation.PodcastId, collab.Invitation.InviterId, collab.Invitation.InviteeId}

	tr, err := c.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	var invitationID pb.ID
	err = tr.QueryRow(query, params...).Scan(&invitationID.Id)
	if err != nil {
		return nil, err
	}

	return &invitationID, nil
}
