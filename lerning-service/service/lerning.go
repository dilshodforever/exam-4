package service

import (
	"context"
	"log"

	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/learning"
	mdb "github.com/dilshodforever/4-oyimtixon-game-service/storage"
)

type LearningService struct {
	stg mdb.InitRoot
	pb.UnimplementedLearningServiceServer
}

func NewLearningService(db mdb.InitRoot) *LearningService {
	return &LearningService{stg:db}
}

func (s *LearningService) GetTopics(ctx context.Context, req *pb.GetTopicsRequest) (*pb.GetTopicsResponse, error) {
	resp, err := s.stg.Game().GetTopics(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) GetTopic(ctx context.Context, req *pb.GetTopicRequest) (*pb.Topic, error) {
	resp, err := s.stg.Game().GetTopic(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) CompleteTopic(ctx context.Context, req *pb.CompleteTopicRequest) (*pb.CompleteTopicResponse, error) {
	resp, err := s.stg.Game().CompleteTopic(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) GetQuiz(ctx context.Context, req *pb.GetQuizRequest) (*pb.Quiz, error) {
	resp, err := s.stg.Game().GetQuiz(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) SubmitQuiz(ctx context.Context, req *pb.SubmitQuizRequest) (*pb.SubmitQuizResponse, error) {
	resp, err := s.stg.Game().SubmitQuiz(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) GetResources(ctx context.Context, req *pb.GetResourcesRequest) (*pb.GetResourcesResponse, error) {
	resp, err := s.stg.Game().GetResources(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) CompleteResource(ctx context.Context, req *pb.CompleteResourceRequest) (*pb.CompleteResourceResponse, error) {
	resp, err := s.stg.Game().CompleteResource(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) GetProgress(ctx context.Context, req *pb.GetProgressRequest) (*pb.ProgressResponse, error) {
	resp, err := s.stg.Game().GetProgress(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) GetRecommendations(ctx context.Context, req *pb.GetRecommendationsRequest) (*pb.GetRecommendationsResponse, error) {
	resp, err := s.stg.Game().GetRecommendations(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) SubmitFeedback(ctx context.Context, req *pb.SubmitFeedbackRequest) (*pb.SubmitFeedbackResponse, error) {
	resp, err := s.stg.Game().SubmitFeedback(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) GetChallenges(ctx context.Context, req *pb.GetChallengesRequest) (*pb.GetChallengesResponse, error) {
	resp, err := s.stg.Game().GetChallenges(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService) SubmitChallengeSolution(ctx context.Context, req *pb.SubmitChallengeSolutionRequest) (*pb.SubmitChallengeSolutionResponse, error) {
	resp, err := s.stg.Game().SubmitChallengeSolution(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}
func (s *LearningService)GetIAstorage(ctx context.Context,req *pb.AistorageRequest) (*pb.AistorageResponse, error) {
	resp, err := s.stg.Game().GetIAstorage(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *LearningService)StartGame(ctx context.Context,req *pb.StartRequest) (*pb.StartResponse, error) {
	resp, err := s.stg.Game().StartGame(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}


