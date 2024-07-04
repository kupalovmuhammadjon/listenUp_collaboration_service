package postgres

// import (
// 	pb "collaboration_service/genproto/collaborations"
// 	"fmt"
// 	"testing"

// 	"github.com/google/uuid"
// )

// func newCollaborationRepoTest() *CollaborationRepo{
// 	db, err := ConnectDB()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return &CollaborationRepo{Db: db}
// }

// func TestCreateInvitation(t *testing.T) {
// 	c := newCollaborationRepoTest()

// 	req := pb.CreateInvite{
// 		PodcastId: "604821bc-c777-4a7e-8f77-4dc783c5a856",
// 		InviterId: "21aa9066-6634-4c3e-ad6e-66e9b00dca70",
// 		InviteeId: "3fac3f55-56ab-46ec-be7f-e596b77134f9",
// 	}

// 	id, err := c.CreateInvitation(&req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = uuid.Parse(id)
// 	if err != nil{
// 		panic(fmt.Errorf("invalid id returned"))
// 	}
	
// }

// func TestCreateCollaboration(t *testing.T) {
// 	c := newCollaborationRepoTest()

// 	req := pb.CreateCollaboration{
// 		Status: "accepted",
// 		InvitationId: "abd93ecf-b6bf-4980-b46c-477d87a15748",
// 		PodcastId: "604821bc-c777-4a7e-8f77-4dc783c5a856",
// 		UserId: "21aa9066-6634-4c3e-ad6e-66e9b00dca70",
// 	}
// 	id, err := c.CreateCollaboration(&req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = uuid.Parse(id)
// 	if err != nil{
// 		panic(fmt.Errorf("invalid id returned"))
// 	}
	
// }

// func TestGetCollaboratorsByPodcastId(t *testing.T){
// 	c := newCollaborationRepoTest()

// 	collaborators, err := c.GetCollaboratorsByPodcastId("604821bc-c777-4a7e-8f77-4dc783c5a856")
// 	if err != nil {
// 		panic(err)
// 	}

// 	if len(collaborators) == 0{
// 		panic(fmt.Errorf("nothing in slice"))
// 	}
// }

// func TestUpdateCollaboratorByPodcastId(t *testing.T){

// 	c := newCollaborationRepoTest()

// 	req := pb.UpdateCollaborator{
// 		Id: "678b8f7b-1090-4660-b487-2bc7a0bf5c63",
// 		PodcastId: "604821bc-c777-4a7e-8f77-4dc783c5a856",
// 		UserId: "21aa9066-6634-4c3e-ad6e-66e9b00dca70",
// 		Role: "owner",
// 	}

// 	err := c.UpdateCollaboratorByPodcastId(&req)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func TestDeleteCollaboratorByPodcastId(t *testing.T){
// 	c := newCollaborationRepoTest()

// 	req := pb.Ids{
// 		PodcastId: "604821bc-c777-4a7e-8f77-4dc783c5a856",
// 		UserId: "21aa9066-6634-4c3e-ad6e-66e9b00dca70",
// 	}

// 	_, err := c.DeleteCollaboratorByPodcastId(&req)
// 	if err != nil {
// 		panic(err)
// 	}
// }


// func TestRespondInvitation(t *testing.T){
// 	c := newCollaborationRepoTest()

// 	req := pb.CreateCollaboration{
// 		Status: "accepted",
// 		InvitationId: "a629f1ef-a4a8-4ab1-bfbb-9f84a72d2748",
// 		PodcastId: "604821bc-c777-4a7e-8f77-4dc783c5a856",
// 		UserId: "21aa9066-6634-4c3e-ad6e-66e9b00dca70",
// 	}
// 	id, err := c.RespondInvitation(&req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = uuid.Parse(id.Id)
// 	if err != nil {
// 		panic(fmt.Errorf("invalid id returned"))
// 	}
	
// 	req = pb.CreateCollaboration{
// 		Status: "declined",
// 		InvitationId: "abd93ecf-b6bf-4980-b46c-477d87a15748",
// 		PodcastId: "604821bc-c777-4a7e-8f77-4dc783c5a856",
// 		UserId: "21aa9066-6634-4c3e-ad6e-66e9b00dca70",
// 	}
// 	id, err = c.RespondInvitation(&req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = uuid.Parse(id.Id)
// 	if err == nil {
// 		panic(fmt.Errorf("valid id returned expected nothing"))
// 	}
// }

// func TestGetPodcastsIdByCollaboratorsId(t *testing.T){
// 	c := newCollaborationRepoTest()
// 	collabIds := []string{"21aa9066-6634-4c3e-ad6e-66e9b00dca70"}
// 	ids, err := c.GetPodcastsIdByCollaboratorsId(&collabIds)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if len(*ids) == 0{
// 		panic(fmt.Errorf("Nothing in the array"))
// 	}
// }

// func TestGetCollaboratorsIdByPodcastsId(t *testing.T){
// 	c := newCollaborationRepoTest()

// 	pIds := []string{"604821bc-c777-4a7e-8f77-4dc783c5a856"}
// 	uIds, err := c.GetCollaboratorsIdByPodcastsId(&pIds)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if len(*uIds) != 1{
// 		panic(fmt.Errorf("expected 1 bu got many"))
// 	}

// }