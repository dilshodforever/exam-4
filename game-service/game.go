package main

import (
	"log"
	"net"

	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/game"
	"github.com/dilshodforever/4-oyimtixon-game-service/service"
	postgres "github.com/dilshodforever/4-oyimtixon-game-service/storage/mongo"
	"github.com/dilshodforever/4-oyimtixon-game-service/kafka"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewMongoConnecti0n()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	liss, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}
	brokers := []string{"10.10.0.175:9092"}

	kcm := kafka.NewKafkaConsumerManager()
	appService := service.NewGameService(db)

	if err := kcm.RegisterConsumer(brokers, "root", "root", kafka.StartLevel(appService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'create-job_application' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}
	s := grpc.NewServer()
	pb.RegisterGameServiceServer(s, service.NewGameService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
