package postgres

import (
	"context"
	"errors"
	"log"

	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/learning"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ls *LearningStorage) GetQuiz(req *pb.GetQuizRequest) (*pb.Quiz, error) {
	coll := ls.db.Collection("topics")
	filter := bson.D{{Key: "quiz.id", Value: req.QuizId}}
	var quiz pb.Quiz
	err := coll.FindOne(context.Background(), filter).Decode(&quiz)
	if err != nil {
		log.Printf("Failed to get quiz: %v", err)
		return nil, err
	}
	return &quiz, nil
}

func (ls *LearningStorage) SubmitQuiz(req *pb.SubmitQuizRequest) (*pb.SubmitQuizResponse, error) {
	coll := ls.db.Collection("topics")
	filter := bson.D{{Key: "quiz.id", Value: req.QuizId}}
	var level pb.Topic
	err := coll.FindOne(context.Background(), filter).Decode(&level)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No level found with id: %v", req.QuizId)
			return nil, errors.New("no rows in result set")
		}
		log.Printf("Failed to decode level: %v", err)
		return nil, err
	}

	var challenge *pb.Quiz
	for _, ch := range level.Quiz {
		if ch.Id == req.QuizId {
			challenge = ch
			break
		}
	}
	if challenge == nil {
		log.Printf("Challenge with id: %v not found in level: %v", req.QuizId, level.Id)
		return nil, errors.New("challenge with id not found in level")
	}

	submitsresult := pb.SubmitQuizResponse{
		TotalQuestions: int32(len(challenge.Questions)),
		CorrectAnswers: make([]*pb.QuizAnswer, len(req.Answers)),
	}

	for i, answer := range req.Answers {
		for _, question := range challenge.Questions {
			if answer.QuestionId == question.Id && answer.SelectedOption == question.CorrectOption {
				submitsresult.XpEarned += 10
				submitsresult.CorrectAnswers[i] = &pb.QuizAnswer{
					QuestionId:     question.Id,
					SelectedOption: answer.SelectedOption,
				}
				break
			}
		}
	}

	correctAnswersCount := 0
	for _, answer := range submitsresult.CorrectAnswers {
		if answer != nil {
			correctAnswersCount++
		}
	}

	if correctAnswersCount == 0 {
		submitsresult.Feedback = "Keep practicing! You can improve."
	} else {
		switch correctAnswersCount {
		case len(req.Answers):
			submitsresult.Feedback = "Excellent! You have a good understanding of the material."
		case len(req.Answers) / 2:
			submitsresult.Feedback = "Nice! You're on the right track."
		default:
			submitsresult.Feedback = "Keep practicing! You can improve."
		}
	}

	_, err = ls.UpdateUserXp(&pb.Update{UserId: req.Userid, Xps: submitsresult.XpEarned})
	if err != nil {
		log.Printf("Failed to update user XP: %v", err)
		return nil, err
	}

	_, err = ls.UpdateComplateds(&pb.CalculateCompleteds{Userid: req.Userid, QuizzesCompleted: int32(correctAnswersCount)})
	if err != nil {
		log.Printf("Failed to update completed quizzes: %v", err)
		return nil, err
	}

	return &submitsresult, nil
}
