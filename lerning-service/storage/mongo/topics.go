package postgres

import (
	"context"
	"log"

	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/learning"
	"go.mongodb.org/mongo-driver/bson"
)

func (ls *LearningStorage) GetTopics(req *pb.GetTopicsRequest) (*pb.GetTopicsResponse, error) {
	coll := ls.db.Collection("topics")
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Failed to get topics: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var topics []*pb.Topic
	for cursor.Next(context.Background()) {
		var topic pb.Topic
		if err := cursor.Decode(&topic); err != nil {
			log.Printf("Failed to decode topic: %v", err)
			return nil, err
		}
		topics = append(topics, &topic)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return &pb.GetTopicsResponse{Topics: topics}, nil
}

func (ls *LearningStorage) GetTopic(req *pb.GetTopicRequest) (*pb.Topic, error) {
	coll := ls.db.Collection("topics")
	filter := bson.D{{Key: "id", Value: req.TopicId}}
	var topic pb.Topic
	err := coll.FindOne(context.Background(), filter).Decode(&topic)
	if err != nil {
		log.Printf("Failed to get topic: %v", err)
		return nil, err
	}
	return &topic, nil
}

func (ls *LearningStorage) CompleteTopic(req *pb.CompleteTopicRequest) (*pb.CompleteTopicResponse, error) {
	coll := ls.db.Collection("topics")
	filter := bson.D{{Key: "id", Value: req.TopicId}}
	var topic pb.Topic
	err := coll.FindOne(context.Background(), filter).Decode(&topic)
	if err != nil {
		log.Printf("Failed to get topic: %v", err)
		return nil, err
	}
	res, err := ls.UpdateUserXp(&pb.Update{UserId: req.Userid, Xps: 50})
	if err != nil {
		log.Printf("Failed to update userxps: %v", err)
		return nil, err
	}
	ls.UpdateComplateds(&pb.CalculateCompleteds{TopicsCompleted: 1})
	return &pb.CompleteTopicResponse{
		Message:  "Topic completed successfully",
		XpEarned: res.XpEarned,
	}, nil
}
