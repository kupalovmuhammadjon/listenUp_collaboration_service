package postgres

import (
	pb "collaboration_service/genproto/comments"
	"database/sql"

	"github.com/google/uuid"
)

type CommentRepo struct {
	Db *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{Db: db}
}

func (c *CommentRepo) CreateCommentByPodcastId(comment *pb.CreateComment) (*pb.ID, error) {
	tx, err := c.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	query := `insert into comments(
		id, podcast_id, user_id, content	
	) values (
		$1, $2, $3, $4
	)`

	newId := uuid.NewString()
	_, err = tx.Exec(query, newId, comment.PodcastId, comment.UserId,
		comment.Content)

	return &pb.ID{Id: newId}, err
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
	    podcast_id = $1`

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
