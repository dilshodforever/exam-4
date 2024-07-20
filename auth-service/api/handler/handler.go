package handler

import (
	pb "github.com/dilshodforever/4-oyimtixon-auth-service/genprotos/auth"
	pbu "github.com/dilshodforever/4-oyimtixon-auth-service/genprotos/user"
	"github.com/dilshodforever/4-oyimtixon-auth-service/kafka"
)

type Handler struct {
	Auth  pb.AuthServiceClient
	User  pbu.UserServiceClient
	Redis InMemoryStorageI
	Kafka kafka.KafkaProducer
}

func NewHandler(auth pb.AuthServiceClient, user pbu.UserServiceClient,
	redis InMemoryStorageI, kafka kafka.KafkaProducer) *Handler {
	return &Handler{
		Auth:  auth,
		User:  user,
		Redis: redis,
		Kafka: kafka,
	}

}
