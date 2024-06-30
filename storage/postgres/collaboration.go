package postgres

import (
	pb "collaboration_service/genproto/collaborations"
	"database/sql"
)

type CollaborationRepo struct {
	Db *sql.DB
}

func NewCollaborationRepo(db *sql.DB) *CollaborationRepo {
	return &CollaborationRepo{Db: db}
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