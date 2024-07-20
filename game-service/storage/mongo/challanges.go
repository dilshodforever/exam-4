package postgres

import (
	"context"
	"errors"
	"log"
	"log/slog"

	"github.com/dilshodforever/4-oyimtixon-game-service/genprotos/game"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (g *GameStorage) GetChallenge(req *game.GetChallengeRequest) (*game.Level, error) {
	levelColl := g.db.Collection("levels")
	var level game.Level
	err := levelColl.FindOne(context.Background(), bson.D{{Key: "challenges.id", Value: req.ChallengeId}}).Decode(&level)
	if err != nil {
		log.Printf("Failed to get level: %v", err)
		return nil, err
	}
	return &game.Level{Challenges: level.Challenges}, nil
}

func (g *GameStorage) SubmitChallenge(req *game.SubmitChallengeRequest) (*game.SubmitChallengeResponse, error) {
	coll := g.db.Collection("levels")
	filter := bson.D{{Key: "challenges.id", Value: req.ChallengeId}}

	var level game.Level
	err := coll.FindOne(context.Background(), filter).Decode(&level)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No level found with id: %v", req.ChallengeId)
			return nil, errors.New("no rows in result set")
		}
		log.Printf("Failed to decode level: %v", err)
		return nil, err
	}

	var challenge *game.Challenge
	for _, ch := range level.Challenges {
		if ch.Id == req.ChallengeId {
			challenge = ch
			break
		}
	}
	if challenge == nil {
		log.Printf("Challenge with id: %v not found in level: %v", req.ChallengeId, level.Levelid)
		return nil, errors.New("challenge with id not found in level")
	}

	var submitsresult game.SubmitChallengeResponse
	submitsresult.TotalQuestions = int32(len(challenge.Questions))

	for i := 0; i < len(req.Answers); i++ {

		for j := 0; j < len(challenge.Questions); j++ {

			if req.Answers[i].SelectedOption == challenge.Questions[j].CorrectOption && req.Answers[i].QuestionId == challenge.Questions[j].Id {
				submitsresult.XpEarned += 10
				submitsresult.CorrectAnswers++
			}
		}
	}
	if submitsresult.XpEarned == 0 {
		submitsresult.Feedback = "Keep practicing! You can improve"
		return &submitsresult, nil
	}
	switch submitsresult.CorrectAnswers {
	case int32(len(req.Answers)):
		submitsresult.Feedback = "Excellent! You have a good understanding of quantum superposition."
	case int32(len(req.Answers)) / 2:
		submitsresult.Feedback = "Nice! You're on the right track."
	default:
		submitsresult.Feedback = "Keep practicing! You can improve."
	}
	_, err = g.UpdateUserXp(&game.CompleteLevelRequest{Userid: req.Userid, Xpearned: submitsresult.XpEarned})
	if err != nil {
		slog.Info(err.Error())
		return nil, err
	}

	return &submitsresult, nil
}
