package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/dilshodforever/4-oyimtixon-game-service/genprotos/game"
	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/learning"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LearningStorage struct {
	db *mongo.Database
}

func NewLearningStorage(db *mongo.Database) *LearningStorage {
	return &LearningStorage{db: db}
}

func (g *LearningStorage) UpdateUserXp(req *pb.Update) (*game.CompleteLevelResponse, error) {
	coll := g.db.Collection("user_levels")
	filter := bson.D{
		{Key: "user_id", Value: req.UserId},
	}
	fmt.Println(req.UserId)
	var user game.Level
	err := coll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to get challenge: %v", err)
		return nil, err
	}
	xps := req.Xps + user.RequiredXp
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "RequiredXp", Value: xps},
		}},
	}
	_, err = coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Failed to complete level: %v", err)
		return nil, err
	}

	return &game.CompleteLevelResponse{XpEarned: xps}, nil
}

func (g *LearningStorage) UpdateComplateds(req *pb.CalculateCompleteds) (*game.CompleteLevelResponse, error) {
	coll := g.db.Collection("Complateds")
	filter := bson.M{
		"user_id": req.Userid,
	}
	fmt.Println(req.Userid)

	update := bson.M{}

	if req.QuizzesCompleted > 0 {
		update["QuizzesCompleted"] = req.QuizzesCompleted
	}
	if req.ResourcesCompleted > 0 {
		update["ResourcesCompleted"] = req.ResourcesCompleted

	}
	if req.TopicsCompleted > 0 {
		update["TopicsCompleted"] = req.TopicsCompleted
	}

	if len(update) == 0 {
		return &game.CompleteLevelResponse{Message: "Nothing to update"}, nil
	}

	// Update the document in the collection
	result, err := coll.UpdateOne(context.Background(), filter, bson.M{"$inc": update})
	if err != nil {
		log.Printf("Failed to update completeds: %v", err)
		return nil, fmt.Errorf("failed to update completeds: %w", err)
	}

	if result.ModifiedCount == 0 {
		return &game.CompleteLevelResponse{Message: "No documents updated"}, nil
	}

	return &game.CompleteLevelResponse{Message: "Success!"}, nil
}

func (ls *LearningStorage) CountQuestions(req []*pb.Topic) int {
	var count int
	for i := 0; i < len(req); i++ {
		count += len(req[i].Quiz)
	}
	return count
}

func (g *LearningStorage) StartGame(req *pb.StartRequest) (*pb.StartResponse, error) {
	coll := g.db.Collection("user_levels")

	document := bson.D{
		{Key: "user_id", Value: req.Userid},
		{Key: "user_xp", Value: 0},
	}

	_, err := coll.InsertOne(context.Background(), document)
	if err != nil {
		log.Printf("Failed to insert user level: %v", err)
		return nil, err
	}

	return &pb.StartResponse{Message: "Success"}, nil
}
