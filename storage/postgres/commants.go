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

	return &pb.ID{Id: newId}, nil
}
