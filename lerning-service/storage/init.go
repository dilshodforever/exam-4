package storage

import (
	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/learning"
)

type InitRoot interface {
	Game() GameStorage
}

type GameStorage interface {
	GetTopics(req *pb.GetTopicsRequest) (*pb.GetTopicsResponse, error)
	GetTopic(req *pb.GetTopicRequest) (*pb.Topic, error)
	CompleteTopic(req *pb.CompleteTopicRequest) (*pb.CompleteTopicResponse, error)
	GetQuiz(req *pb.GetQuizRequest) (*pb.Quiz, error)
	SubmitQuiz(req *pb.SubmitQuizRequest) (*pb.SubmitQuizResponse, error)
	GetResources(req *pb.GetResourcesRequest) (*pb.GetResourcesResponse, error)
	CompleteResource(req *pb.CompleteResourceRequest) (*pb.CompleteResourceResponse, error)
	GetProgress(req *pb.GetProgressRequest) (*pb.ProgressResponse, error)
	GetRecommendations(req *pb.GetRecommendationsRequest) (*pb.GetRecommendationsResponse, error)
	SubmitFeedback(req *pb.SubmitFeedbackRequest) (*pb.SubmitFeedbackResponse, error)
	GetChallenges(req *pb.GetChallengesRequest) (*pb.GetChallengesResponse, error)
	SubmitChallengeSolution(req *pb.SubmitChallengeSolutionRequest) (*pb.SubmitChallengeSolutionResponse, error)
	GetIAstorage(req *pb.AistorageRequest) (*pb.AistorageResponse, error)
	StartGame(req *pb.StartRequest) (*pb.StartResponse, error)
}
