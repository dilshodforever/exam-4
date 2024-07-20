package postgres

import (
	"context"
	"log"

	"github.com/dilshodforever/4-oyimtixon-game-service/genprotos/game"
	"go.mongodb.org/mongo-driver/bson"
)

func (g *GameStorage) GetLevels(req *game.GetLevelsRequest) (*game.GetLevelsResponse, error) {
	coll := g.db.Collection("levels")
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Failed to get levels: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var levels []*game.Level
	for cursor.Next(context.Background()) {
		var level game.Level
		if err := cursor.Decode(&level); err != nil {
			log.Printf("Failed to decode level: %v", err)
			return nil, err
		}
		levels = append(levels, &level)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return &game.GetLevelsResponse{Levels: levels}, nil
}

func (g *GameStorage) StartLevel(req *game.StartLevelRequest) (*game.StartLevelResponse, error) {
	coll := g.db.Collection("user_levels")
	_, err := coll.InsertOne(context.Background(), bson.D{
		{Key: "user_id", Value: req.Userid},
		{Key: "level_id", Value: "7a4b0c7c-54b5-4a9f-9b9b-14e8cf78b8b5"},
		{Key: "status", Value: "started"},
		{Key: "user_xp", Value: 0},
	})
	if err != nil {
		log.Printf("Failed to start level: %v", err)
		return nil, err
	}

	return &game.StartLevelResponse{
		Message: "Level started successfully",
	}, nil
}

func (g *GameStorage) CompleteLevel(req *game.CompleteLevelRequest) (*game.CompleteLevelResponse, error) {
	coll := g.db.Collection("user_levels")
	filter := bson.D{
		{Key: "user_id", Value: req.Userid},
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "completed"},
			{Key: "level_id", Value: req.LevelId},
		}},
	}
	_, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Failed to complete level: %v", err)
		return nil, err
	}

	return &game.CompleteLevelResponse{
		Message:          "Level completed successfully",
		XpEarned:         req.Xpearned,
		NewLevelUnlocked: req.LevelId,
	}, nil
}
