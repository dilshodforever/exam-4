package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/dilshodforever/4-oyimtixon-api-gatway/genprotos/game"
)

// GetLevels handles retrieving game levels
// @Summary      Get Levels
// @Description  Get details of the game levels
// @Tags         Game
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} pb.GetLevelsResponse "Game levels details"
// @Failure      400 {string} string "Missing required query parameter"
// @Failure      500 {string} string "Error while fetching levels"
// @Router       /game/get_levels [get]
func (h *Handler) GetLevels(ctx *gin.Context) {
	id := ctx.Query("id")

	req := &pb.GetLevelsRequest{
		Id: id,
	}

	res, err := h.Game.GetLevels(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// StartLevel handles starting a game level
// @Summary      Start Level
// @Description  Start a new game level
// @Tags         Game
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.StartLevelRequest true "Start level request"
// @Success      200 {object} pb.StartLevelResponse "Level started successfully"
// @Failure      500 {string} string "Error while starting level"
// @Router       /game/start_level [post]
func (h *Handler) StartLevel(ctx *gin.Context) {
	req := &pb.StartLevelRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.Game.StartLevel(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// CompleteLevel handles completing a game level
// @Summary      Complete Level
// @Description  Complete an ongoing game level
// @Tags         Game
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.CompleteLevelRequest true "Complete level request"
// @Success      200 {object} pb.CompleteLevelResponse "Level completed successfully"
// @Failure      500 {string} string "Error while completing level"
// @Router       /game/complete_level [post]
func (h *Handler) CompleteLevel(ctx *gin.Context) {
	req := &pb.CompleteLevelRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.Game.CompleteLevel(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetChallenge handles retrieving a challenge by ID
// @Summary      Get Challenge
// @Description  Retrieve details of a challenge by ID
// @Tags         Game
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id query string true "Challenge ID"
// @Success      200 {object} pb.Challenge "Challenge details"
// @Failure      400 {string} string "Missing required query parameter"
// @Failure      500 {string} string "Error while fetching challenge"
// @Router       /game/get_challenge [get]
func (h *Handler) GetChallenge(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		ctx.JSON(400, gin.H{"error": "Missing required query parameter: id"})
		return
	}

	req := &pb.GetChallengeRequest{
		ChallengeId: id,
	}

	res, err := h.Game.GetChallenge(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// SubmitChallenge handles submitting a challenge response
// @Summary      Submit Challenge
// @Description  Submit a response to a challenge
// @Tags         Game
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.SubmitChallengeRequest true "Submit challenge request"
// @Success      200 {object} pb.SubmitChallengeResponse "Challenge submitted successfully"
// @Failure      500 {string} string "Error while submitting challenge"
// @Router       /game/submit_challenge [post]
func (h *Handler) SubmitChallenge(ctx *gin.Context) {
	req := &pb.SubmitChallengeRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.Game.SubmitChallenge(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}




// GetLeaderboard handles retrieving the game leaderboard
// @Summary      Get Leaderboard
// @Description  Retrieve leaderboard based on criteria
// @Tags         Game
// @Accept       json
// @Produce      json
// @Security     BearerAuthz
// @Success      200 {object} pb.LeaderboardResponse "Leaderboard details"
// @Failure      400 {string} string "Missing required query parameter"
// @Failure      500 {string} string "Error while fetching leaderboard"
// @Router       /game/get_leaderboard [get]
func (h *Handler) GetLeaderboard(ctx *gin.Context) {

	req := &pb.GetLeaderboardRequest{}

	res, err := h.Game.GetLeaderboard(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetAchievements handles retrieving achievements
// @Summary      Get Achievements
// @Description  Retrieve achievements based on criteria
// @Tags         Game
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        player_id query string true "Player ID"
// @Success      200 {object} pb.AchievementsResponse "Achievements list"
// @Failure      400 {string} string "Missing required query parameter"
// @Failure      500 {string} string "Error while fetching achievements"
// @Router       /game/get_achievements [get]
func (h *Handler) GetAchievements(ctx *gin.Context) {
	playerID := ctx.Query("player_id")

	if playerID == "" {
		ctx.JSON(400, gin.H{"error": "Missing required query parameter: player_id"})
		return
	}

	req := &pb.GetAchievementsRequest{
		Token: playerID,
	}

	res, err := h.Game.GetAchievements(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}
