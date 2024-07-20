package main

import (
	"fmt"
	"log"

	"github.com/dilshodforever/4-oyimtixon-api-gatway/api"
	"github.com/dilshodforever/4-oyimtixon-api-gatway/api/handler"
	pb "github.com/dilshodforever/4-oyimtixon-api-gatway/genprotos/auth"
	pbg "github.com/dilshodforever/4-oyimtixon-api-gatway/genprotos/game"
	pbl "github.com/dilshodforever/4-oyimtixon-api-gatway/genprotos/learning"
	pbu "github.com/dilshodforever/4-oyimtixon-api-gatway/genprotos/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	UserConn, err := grpc.NewClient(fmt.Sprintf("auth-service%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()

	GameConn, err := grpc.NewClient(fmt.Sprintf("gameservice%s", ":8087"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer GameConn.Close()
	LearningConn, err := grpc.NewClient(fmt.Sprintf("learning-service%s", ":8088"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer LearningConn.Close()

	auth := pb.NewAuthServiceClient(UserConn)
	user := pbu.NewUserServiceClient(UserConn)
	game := pbg.NewGameServiceClient(GameConn)
	lerning := pbl.NewLearningServiceClient(LearningConn)
	h := handler.NewHandler(game, auth, user, lerning)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}
