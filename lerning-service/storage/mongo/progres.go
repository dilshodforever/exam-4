package postgres

import (
	"context"
	"errors"
	"log"

	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/learning"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ls *LearningStorage) GetProgress(req *pb.GetProgressRequest) (*pb.ProgressResponse, error) {
	coll := ls.db.Collection("Complateds")
	filter := bson.D{{Key: "user_id", Value: req.Userid}}

	var completedData pb.CalculateCompleteds
	err := coll.FindOne(context.Background(), filter).Decode(&completedData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No progress found for user_id: %v", req.Userid)
			return nil, errors.New("no progress data found")
		}
		log.Printf("Failed to get progress: %v", err)
		return nil, err
	}

	res, err := ls.GetTopics(&pb.GetTopicsRequest{})
	if err != nil {
		log.Printf("Failed to get topics: %v", err)
		return nil, err
	}
	rest, err := ls.GetResources(&pb.GetResourcesRequest{})
	if err != nil {
		log.Printf("Failed to get resorce: %v", err)
		return nil, err
	}
	totalTopics := len(res.Topics)
	totalQuizzes := ls.CountQuestions(res.Topics)
	totalResources := len(rest.Resources)

	overallProgress := float32(completedData.TopicsCompleted+completedData.QuizzesCompleted+completedData.ResourcesCompleted) /
		float32(totalTopics+totalQuizzes+totalResources) * 100

	progress := &pb.ProgressResponse{
		TopicsCompleted:    completedData.TopicsCompleted,
		TotalTopics:        int32(totalTopics),
		QuizzesCompleted:   completedData.QuizzesCompleted,
		TotalQuizzes:       int32(totalQuizzes),
		ResourcesCompleted: completedData.ResourcesCompleted,
		TotalResources:     int32(totalResources),
		OverallProgress:    overallProgress,
	}

	return progress, nil
}

func (ls *LearningStorage) GetRecommendations(req *pb.GetRecommendationsRequest) (*pb.GetRecommendationsResponse, error) {
	coll := ls.db.Collection("recommendations")
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Failed to get topics: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	recommendations := pb.GetRecommendationsResponse{}
	for cursor.Next(context.Background()) {
		var topic pb.Topics
		if err := cursor.Decode(&topic); err != nil {
			log.Printf("Failed to decode topic: %v", err)
			return nil, err
		}
		recommendations.Recommendations = append(recommendations.Recommendations, &topic)
	}
	return &recommendations, nil
}

func (ls *LearningStorage) SubmitFeedback(req *pb.SubmitFeedbackRequest) (*pb.SubmitFeedbackResponse, error) {
	coll := ls.db.Collection("feedbacks")
	documents := []interface{}{
		bson.D{
			{Key: "Userid", Value: req.Userid},
			{Key: "TopicId", Value: req.TopicId},
			{Key: "rating", Value: req.Rating},
			{Key: "comment", Value: req.Comment},
		},
	}

	_, err := coll.InsertMany(context.Background(), documents)
	if err != nil {
		log.Printf("Failed to insert feedback: %v", err)
		return nil, err
	}

	res, err := ls.UpdateUserXp(&pb.Update{UserId: req.Userid, Xps: 10})
	if err != nil {
		log.Printf("Failed to update xps: %v", err)
		return nil, err
	}
	return &pb.SubmitFeedbackResponse{
		Message:  "Feedback submitted successfully",
		XpEarned: res.XpEarned,
	}, nil
}

func (ls *LearningStorage) GetChallenges(req *pb.GetChallengesRequest) (*pb.GetChallengesResponse, error) {
	coll := ls.db.Collection("challenges")
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Failed to get challenges: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var challenges []*pb.Challenge
	for cursor.Next(context.Background()) {
		var challenge pb.Challenge
		if err := cursor.Decode(&challenge); err != nil {
			log.Printf("Failed to decode challenge: %v", err)
			return nil, err
		}
		challenges = append(challenges, &challenge)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return &pb.GetChallengesResponse{Challenges: challenges}, nil
}
