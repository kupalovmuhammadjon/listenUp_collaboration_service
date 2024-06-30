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

func (c *CollaborationRepo) GetCollaboratorsByPodcastId(PodcastId string) (*[]pb.CollaboratorToGet, error) {
	query := `select user_id, role, joined_at from collaborations
	where podcast_id = $1`

	callabrators := []pb.CollaboratorToGet{}
	rows, err := c.Db.Query(query, PodcastId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := pb.CollaboratorToGet{}
		err := rows.Scan(&c.UserId, &c.Role, &c.JoinedAt)
		if err != nil {
			return nil, err
		}
		callabrators = append(callabrators, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &callabrators, nil
}
