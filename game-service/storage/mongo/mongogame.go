package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/dilshodforever/4-oyimtixon-game-service/genprotos/game"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LeaderboardEntry struct {
	UserID  string `bson:"user_id" json:"user_id"`
	LevelID string `bson:"level_id" json:"level_id"`
	UserXP  int32  `bson:"user_xp" json:"user_xp"`
}

type UserLevel struct {
	UserID  string `bson:"user_id"`
	UserXP  int32  `bson:"user_xp"`
	Levelid string `bson:"level_id"`
}

type GameStorage struct {
	db *mongo.Database
}

func NewGameStorage(db *mongo.Database) *GameStorage {
	return &GameStorage{db: db}
}
func (g *GameStorage) GetLeaderboard(req *game.GetLeaderboardRequest) (*game.LeaderboardResponse, error) {
	coll := g.db.Collection("user_levels")

	findOptions := options.Find().SetSort(bson.D{{Key: "user_xp", Value: -1}})

	
	cursor, err := coll.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		log.Printf("Failed to get leaderboard: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var leaderboard []*game.LeaderboardEntry
	for cursor.Next(context.Background()) {
		var entry LeaderboardEntry
		if err := cursor.Decode(&entry); err != nil {
			log.Printf("Failed to decode leaderboard entry: %v", err)
			return nil, err
		}

		lb := game.LeaderboardEntry{
			UserId:  entry.UserID,
			LevelId: entry.LevelID,
			UserXp:  entry.UserXP,
			Rank:    entry.UserXP * 2, 
		}
		leaderboard = append(leaderboard, &lb)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return &game.LeaderboardResponse{
		Leaderboard: leaderboard,
	}, nil
}


func (g *GameStorage) GetAchievements(req *game.GetAchievementsRequest) (*game.AchievementsResponse, error) {
	coll := g.db.Collection("achievements")
	filter := bson.D{{Key: "user_id", Value: req.Token}}
	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		log.Printf("Failed to get achievements: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var achievements []*game.Achievement
	for cursor.Next(context.Background()) {
		var achievement game.Achievement
		if err := cursor.Decode(&achievement); err != nil {
			log.Printf("Failed to decode achievement: %v", err)
			return nil, err
		}
		achievements = append(achievements, &achievement)
	}

	return &game.AchievementsResponse{Achievements: achievements}, nil
}

func (g *GameStorage) CheckLevels(req *game.Cheak) (*game.CHeakResult, error) {
	coll := g.db.Collection("levels")
	filter := bson.D{{Key: "Levelid", Value: req.Levelid}}

	var level game.Level
	err := coll.FindOne(context.Background(), filter).Decode(&level)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No level found with id: %v", req.Levelid)
			return &game.CHeakResult{Result: false}, nil
		}
		log.Printf("Failed to decode level: %v", err)
		return &game.CHeakResult{Result: false}, err
	}

	if level.RequiredXp < req.Userxp {
		cid := level.Cid + 1
		newFilter := bson.D{{Key: "cid", Value: cid}}
		var newLevel game.Level
		err := coll.FindOne(context.Background(), newFilter).Decode(&newLevel)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				log.Printf("No level found with required xp: %v", err)
				return &game.CHeakResult{Result: false}, nil
			}
			log.Printf("Failed to decode new level: %v", err)
			return &game.CHeakResult{Result: false}, err
		}
		return &game.CHeakResult{Result: true, Levelid: newLevel.Levelid, Xpearned: newLevel.RequiredXp}, nil
	}
	return &game.CHeakResult{Result: false}, nil
}

func (g *GameStorage) UpdateUserXp(req *game.CompleteLevelRequest) (*game.CompleteLevelResponse, error) {
	coll := g.db.Collection("user_levels")
	filter := bson.D{
		{Key: "user_id", Value: req.Userid},
	}
	var user UserLevel
	err := coll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Printf("Failed to get challenge: %v", err)
		return nil, err
	}
	xps := req.Xpearned + user.UserXP
	res, err := g.CheckLevels(&game.Cheak{Levelid: user.Levelid, Userxp: xps})
	if err != nil {
		log.Printf("Failed to cheacklevels: %v", err)
		return nil, err
	}
	if res.Result {
		_, err = g.CompleteLevel(&game.CompleteLevelRequest{Userid: req.Userid, LevelId: res.Levelid, Xpearned: res.Xpearned})
		fmt.Println(11111)
		if err != nil {
			log.Printf("Failed to complatelevel: %v", err)
			return nil, err
		}
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "user_xp", Value: xps},
		}},
	}
	_, err = coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Failed to complete level: %v", err)
		return nil, err
	}

	return &game.CompleteLevelResponse{}, nil
}
