package postgres

import (
	pb "collaboration_service/genproto/comments"
	"database/sql"
)

type CommentRepo struct {
	Db *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{Db: db}
}

func (c *CommentRepo) GetCommentsByPodcastId(id *pb.ID) (*[]pb.CommentInfo, error) {
	query := `
	select
		user_id,
    	content,
    	created_at,
    	updated_at
	from 
	    comments
	where
	    podcast_id = $1
`
	comments := []pb.CommentInfo{}
	rows, err := c.Db.Query(query, id.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment pb.CommentInfo
		err = rows.Scan(&comment.UserId, &comment.Content, comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return &comments, nil
}
