package postgres

import (
	pb "collaboration_service/genproto/comments"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func newCommentRepoTest() *CommentRepo {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	return &CommentRepo{Db: db}
}

func TestCreateCommentByPodcastId(t *testing.T) {
	c := newCommentRepoTest()

	req := pb.CreateComment{
		PodcastId: "604821bc-c777-4a7e-8f77-4dc783c5a856",
		UserId: "21aa9066-6634-4c3e-ad6e-66e9b00dca70",
		Content: "Vashshe zor",
	}

	id, err := c.CreateCommentByPodcastId(&req)
	if err != nil {
		panic(err)
	}

	_, err = uuid.Parse(id.Id)
	if err != nil {
		panic(err)
	}

}

func TestCreateEpisodeComment(t *testing.T) {
	c := newCommentRepoTest()

	req := pb.EpisodeComment{
		EpisodeId: "604821bc-c777-4a7e-8f77-4dc783c5a856",
		UserId: "21aa9066-6634-4c3e-ad6e-66e9b00dca70",
		Content: "Vashshe zor",
	}

	id, err := c.CreateEpisodeComment(&req)
	if err != nil {
		panic(err)
	}

	_, err = uuid.Parse(id)
	if err != nil {
		panic(err)
	}

}

func TestGetCommentsByPodcastId(t *testing.T){
	c := newCommentRepoTest()

	req := pb.ID{
		Id: "604821bc-c777-4a7e-8f77-4dc783c5a856",
	}
	comments, err := c.GetCommentsByPodcastId(&req)
	if err != nil {
		panic(err)
	}
	if len(comments) == 0{
		panic(fmt.Errorf("expected more got 0"))
	}
}

func TestGetCommentsByEpisodeId(t *testing.T){
	c := newCommentRepoTest()

	req := pb.ID{
		Id: "604821bc-c777-4a7e-8f77-4dc783c5a856",
	}
	comments, err := c.GetCommentsByEpisodeId(&req)
	if err != nil {
		panic(err)
	}
	if len(comments) == 0{
		panic(fmt.Errorf("expected more got 0"))
	}
}