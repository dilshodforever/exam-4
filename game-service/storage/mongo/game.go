package postgres

// import (
// 	"database/sql"
// 	"log"

// 	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/game"
// )

// type GameStorage struct {
// 	db *sql.DB
// }

// func NewGameStorage(db *sql.DB) *GameStorage {
// 	return &GameStorage{db: db}
// }

// func (g *GameStorage) GetLevels(req *pb.GetLevelsRequest) (*pb.GetLevelsResponse, error) {
// 	query := `SELECT id, name, description, required_xp, completed FROM levels WHERE`
// 	rows, err := g.db.Query(query)
// 	if err != nil {
// 		log.Printf("Failed to get levels: %v", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var levels []*pb.Level
// 	for rows.Next() {
// 		var level pb.Level
// 		if err := rows.Scan(&level.Id, &level.Name, &level.Description, &level.RequiredXp, &level.Completed); err != nil {
// 			log.Printf("Failed to scan level: %v", err)
// 			return nil, err
// 		}
// 		levels = append(levels, &level)
// 	}

// 	return &pb.GetLevelsResponse{Levels: levels}, nil
// }




// func (g *GameStorage) StartLevel(req *pb.StartLevelRequest) (*pb.StartLevelResponse, error) {
// 	query := `INSERT INTO user_levels (user_id, level_id, status) VALUES ($1, $2, 'started')`
// 	_, err := g.db.Exec(query, req.Token, req.LevelId)
// 	if err != nil {
// 		log.Printf("Failed to start level: %v", err)
// 		return nil, err
// 	}

// 	// Assuming the first challenge is determined by some logic
// 	firstChallengeId := "first-challenge-id"

// 	return &pb.StartLevelResponse{
// 		Message:         "Level started successfully",
// 		FirstChallengeId: firstChallengeId,
// 	}, nil
// }

// func (g *GameStorage) CompleteLevel(req *pb.CompleteLevelRequest) (*pb.CompleteLevelResponse, error) {
// 	query := `UPDATE user_levels SET status='completed' WHERE user_id=$1 AND level_id=$2`
// 	_, err := g.db.Exec(query, req.Token, req.LevelId)
// 	if err != nil {
// 		log.Printf("Failed to complete level: %v", err)
// 		return nil, err
// 	}

// 	// Assuming XP earned and new level unlocked are determined by some logic
// 	var xpEarned int32 = 100
// 	newLevelUnlocked := "next-level-id"

// 	return &pb.CompleteLevelResponse{
// 		Message:           "Level completed successfully",
// 		XpEarned:          xpEarned,
// 		NewLevelUnlocked:  newLevelUnlocked,
// 	}, nil
// }

// func (g *GameStorage) GetChallenge(req *pb.GetChallengeRequest) (*pb.Level, error) {
// 	query := `SELECT ch.id, ch.name, ch.type, ch.completed, u.name FROM challenges 
// 	join user_levels u  on ch.id = u.challangeid WHERE ch.id=$1`
// 	var levels pb.Level
// 	row,err:= g.db.Query(query, req.ChallengeId)
// 	if err!=nil{
// 		return nil, err
// 	}
// 	for row.Next(){
// 	var level pb.Level
// 	var challenge pb.Challenge
// 	if err := row.Scan(&challenge.Id, &challenge.Name, &challenge.Type, &challenge.Completed,
// 		&level.Name,&level.Description, &level.RequiredXp); err != nil {
// 		log.Printf("Failed to get challenge: %v", err)
// 		return nil, err
// 	}
// 	levels.Challenges=append(levels.Challenges, &challenge)

// 	}
	
// 	return &levels, nil
// }

// func (g *GameStorage) SubmitChallenge(req *pb.SubmitChallengeRequest) (*pb.SubmitChallengeResponse, error) {
// 	// Assuming logic to check answers and calculate results
// 	var correctAnswers int32= 3
// 	var totalQuestions int32= 5
// 	var xpEarned int32 = 50
// 	feedback := "Well done!"

// 	return &pb.SubmitChallengeResponse{
// 		CorrectAnswers: correctAnswers,
// 		TotalQuestions: totalQuestions,
// 		XpEarned:       xpEarned,
// 		Feedback:       feedback,
// 	}, nil
// }


// func (g *GameStorage) GetLeaderboard(req *pb.GetLeaderboardRequest) (*pb.LeaderboardResponse, error) {
// 	query := `SELECT rank, username, level, xp FROM leaderboard`
// 	rows, err := g.db.Query(query)
// 	if err != nil {
// 		log.Printf("Failed to get leaderboard: %v", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var leaderboard []*pb.LeaderboardEntry
// 	for rows.Next() {
// 		var entry pb.LeaderboardEntry
// 		if err := rows.Scan(&entry.Rank, &entry.Username, &entry.Level, &entry.Xp); err != nil {
// 			log.Printf("Failed to scan leaderboard entry: %v", err)
// 			return nil, err
// 		}
// 		leaderboard = append(leaderboard, &entry)
// 	}

// 	// Assuming user rank is determined by some logic
// 	var userRank int32 =10

// 	return &pb.LeaderboardResponse{
// 		Leaderboard: leaderboard,
// 		UserRank:    userRank,
// 	}, nil
// }

// func (g *GameStorage) GetAchievements(req *pb.GetAchievementsRequest) (*pb.AchievementsResponse, error) {
// 	query := `SELECT id, name, description, earned_at FROM achievements WHERE user_id=$1`
// 	rows, err := g.db.Query(query, req.Token)
// 	if err != nil {
// 		log.Printf("Failed to get achievements: %v", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var achievements []*pb.Achievement
// 	for rows.Next() {
// 		var achievement pb.Achievement
// 		if err := rows.Scan(&achievement.Id, &achievement.Name, &achievement.Description, &achievement.EarnedAt); err != nil {
// 			log.Printf("Failed to scan achievement: %v", err)
// 			return nil, err
// 		}
// 		achievements = append(achievements, &achievement)
// 	}

// 	return &pb.AchievementsResponse{Achievements: achievements}, nil
// }
