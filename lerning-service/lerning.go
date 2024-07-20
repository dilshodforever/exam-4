package main

import (
	"log"
	"net"
	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/learning"
	"github.com/dilshodforever/4-oyimtixon-game-service/service"
	postgres "github.com/dilshodforever/4-oyimtixon-game-service/storage/mongo"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewMongoConnecti0n()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	liss, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterLearningServiceServer(s, service.NewLearningService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
}
