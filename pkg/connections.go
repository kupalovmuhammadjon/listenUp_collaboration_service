package pkg

import (
	"collaboration_service/config"
	pbp "collaboration_service/genproto/podcasts"
	pbe "collaboration_service/genproto/episodes"
	pbu "collaboration_service/genproto/user"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateUserManagementClient(cfg *config.Config) pbu.UserManagementClient {
	conn, err := grpc.NewClient(cfg.USER_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
	}
	defer conn.Close()

	u := pbu.NewUserManagementClient(conn)
	return u
}

func NewPodcastsClient(cfg *config.Config) pbp.PodcastsClient {
	conn, err := grpc.NewClient(cfg.PODCAST_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting podcast service ", err)
	}
	a := pbp.NewPodcastsClient(conn)

	return a
}

func NewEpisodesClient(cfg *config.Config) pbe.EpisodesServiceClient {
	conn, err := grpc.NewClient(cfg.PODCAST_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("error while connecting podcast service ", err)
	}
	a := pbe.NewEpisodesServiceClient(conn)

	return a
}
