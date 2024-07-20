package handler

import (
	pba "github.com/dilshodforever/4-oyimtixon-api-gatway/genprotos/auth"
	pb "github.com/dilshodforever/4-oyimtixon-api-gatway/genprotos/game"
	pbu "github.com/dilshodforever/4-oyimtixon-api-gatway/genprotos/user"
	"github.com/dilshodforever/4-oyimtixon-api-gatway/kafka"
    pbl"github.com/dilshodforever/4-oyimtixon-api-gatway/genprotos/learning"
)

type Handler struct {
	Game pb.GameServiceClient
    Auth pba.AuthServiceClient
    User pbu.UserServiceClient
    Kafka kafka.KafkaProducer
    LearningService pbl.LearningServiceClient
}

func NewHandler(game pb.GameServiceClient, 
    auth pba.AuthServiceClient, 
    user pbu.UserServiceClient, 
    learningService pbl.LearningServiceClient) *Handler {
	return &Handler{
		Game: game,
        Auth: auth,
        User: user,
        LearningService: learningService,
	}
}
