// Package handler provides HTTP handlers for the learning service.
package handler

import (
	"log"
	"log/slog"

	"github.com/dilshodforever/4-oyimtixon-api-gatway/api/token"
	"github.com/dilshodforever/4-oyimtixon-api-gatway/config"

	pb "github.com/dilshodforever/4-oyimtixon-api-gatway/genprotos/learning"
	pbc "github.com/dilshodforever/4-oyimtixon-api-gatway/genprotos/user"
	"github.com/gin-gonic/gin"
)

// @Summary Get all topics
// @Description Retrieves all topics available
// @Tags Learning
// @Accept json
// @Produce json
// @Success 200 {object} pb.GetTopicsResponse "Topics retrieved successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /learning/topics [get]
func (h *Handler) GetTopics(ctx *gin.Context) {
	req := &pb.GetTopicsRequest{}
	res, err := h.LearningService.GetTopics(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(200, res)
}

// @Summary Get a topic by ID
// @Description Get details of a specific topic
// @Tags Learning
// @Accept json
// @Produce json
// @Param topic_id path string true "Topic ID"
// @Success 200 {object} pb.Topic "Topic details"
// @Failure 404 {string} string "Topic not found"
// @Router /learning/topic/{topic_id} [get]
func (h *Handler) GetTopic(ctx *gin.Context) {
	id := ctx.Param("topic_id")
	req := &pb.GetTopicRequest{TopicId: id}
	res, err := h.LearningService.GetTopic(ctx, req)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Topic not found"})
		return
	}
	ctx.JSON(200, res)
}

// @Summary Complete a topic
// @Description Mark a topic as completed for a specific user
// @Tags Learning
// @Accept json
// @Produce json
// @Param body body pb.CompleteTopicRequest true "Complete Topic Request"
// @Success 200 {object} pb.CompleteTopicResponse "Topic completed successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Topic not found"
// @Router /learning/topic/complete [post]
func (h *Handler) CompleteTopic(ctx *gin.Context) {
	var req pb.CompleteTopicRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	res, err := h.LearningService.CompleteTopic(ctx, &req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	_, err = h.User.UpdateXps(ctx, &pbc.UpdateXpRequest{Userid: req.Userid, Xp: res.XpEarned})
	if err != nil {
		slog.Info(err.Error())
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	ctx.JSON(200, res)
}

// @Summary Submit a quiz
// @Description Submit answers for a quiz
// @Tags Learning
// @Accept json
// @Produce json
// @Param body body pb.SubmitQuizRequest true "Submit Quiz Request"
// @Success 200 {object} pb.SubmitQuizResponse "Quiz submitted successfully"
// @Failure 400 {string} string "Bad request"
// @Router /learning/quiz/submit [post]
func (h *Handler) SubmitQuiz(ctx *gin.Context) {
	var req pb.SubmitQuizRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	res, err := h.LearningService.SubmitQuiz(ctx, &req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	_, err = h.User.UpdateXps(ctx, &pbc.UpdateXpRequest{Userid: req.Userid, Xp: res.XpEarned})
	if err != nil {
		slog.Info(err.Error())
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	ctx.JSON(200, res)
}

// @Summary Get all resources
// @Description Retrieves all resources available
// @Tags Learning
// @Accept json
// @Produce json
// @Success 200 {object} pb.GetResourcesResponse "Resources retrieved successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /learning/resources [get]
func (h *Handler) GetResources(ctx *gin.Context) {
	req := &pb.GetResourcesRequest{}
	res, err := h.LearningService.GetResources(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(200, res)
}

// @Summary Complete a resource
// @Description Mark a resource as completed for a specific user
// @Tags Learning
// @Accept json
// @Produce json
// @Param body body pb.CompleteResourceRequest true "Complete Resource Request"
// @Success 200 {object} pb.CompleteResourceResponse "Resource completed successfully"
// @Failure 400 {string} string "Bad request"
// @Router /learning/resource/complete [post]
func (h *Handler) CompleteResource(ctx *gin.Context) {
	var req pb.CompleteResourceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	res, err := h.LearningService.CompleteResource(ctx, &req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	_, err = h.User.UpdateXps(ctx, &pbc.UpdateXpRequest{Userid: req.UserId, Xp: res.XpEarned})
	if err != nil {
		slog.Info(err.Error())
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	ctx.JSON(200, res)
}

// @Summary Get progress for a user
// @Description Retrieves progress details for a specific user
// @Tags Learning
// @Accept json
// @Produce json
// @Success 200 {object} pb.ProgressResponse "Progress details"
// @Failure 400 {string} string "Bad request"
// @Router /learning/progress [get]
func (h *Handler) GetProgress(ctx *gin.Context) {
	var req pb.GetProgressRequest
	req.Userid="1a0df0c6-41df-479e-ad2d-4339a05fd53a"
	res, err := h.LearningService.GetProgress(ctx, &req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	ctx.JSON(200, res)
}

// @Summary Get recommendations
// @Description Retrieves topic recommendations
// @Tags Learning
// @Accept json
// @Produce json
// @Success 200 {object} pb.GetRecommendationsResponse "Recommendations retrieved successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /learning/recommendations [get]
func (h *Handler) GetRecommendations(ctx *gin.Context) {
	req := &pb.GetRecommendationsRequest{}
	res, err := h.LearningService.GetRecommendations(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(200, res)
}

// @Summary Submit feedback for a topic
// @Description Submit feedback and earn XP for a specific topic
// @Tags Learning
// @Accept json
// @Produce json
// @Param body body pb.SubmitFeedbackRequest true "Submit Feedback Request"
// @Success 200 {object} pb.SubmitFeedbackResponse "Feedback submitted successfully"
// @Failure 400 {string} string "Bad request"
// @Router /learning/feedback/submit [post]

func (h *Handler) SubmitFeedback(ctx *gin.Context) {
	var req pb.SubmitFeedbackRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	res, err := h.LearningService.SubmitFeedback(ctx, &req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	_, err = h.User.UpdateXps(ctx, &pbc.UpdateXpRequest{Userid: req.Userid, Xp: res.XpEarned})
	if err != nil {
		log.Fatalf("error while updating user xp: %v", err)
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	ctx.JSON(200, res)
}

// @Summary Get all challenges
// @Description Retrieves all challenges available
// @Tags Learning
// @Accept json
// @Produce json
// @Success 200 {object} pb.GetChallengesResponse "Challenges retrieved successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /learning/challenges [get]
func (h *Handler) GetChallenges(ctx *gin.Context) {
	req := &pb.GetChallengesRequest{}
	res, err := h.LearningService.GetChallenges(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(200, res)
}

// @Summary Submit solution for a challenge
// @Description Submit a solution for a challenge
// @Tags Learning
// @Accept json
// @Produce json
// @Param body body pb.SubmitChallengeSolutionRequest true "Submit Challenge Solution Request"
// @Success 200 {object} pb.SubmitChallengeSolutionResponse "Solution submitted successfully"
// @Failure 400 {string} string "Bad request"
// @Router /learning/challenge/solution/submit [post]
func (h *Handler) SubmitChallengeSolution(ctx *gin.Context) {
	var req pb.SubmitChallengeSolutionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	// jwtToken := ctx.Request.Header.Get("Authourization")
	// claims, err:=token.ExtractClaim(&config.Config{TokenKey: config.Load().TokenKey},jwtToken)
	// if err != nil {
	// 	ctx.JSON(400, err)
	// 	return
	// }
	// req.Userid=claims["username"].(string)
	res, err := h.LearningService.SubmitChallengeSolution(ctx, &req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	ctx.JSON(200, res)
}

// @Summary Get storage from IA
// @Description Get AI storage information
// @Tags Learning
// @Accept json
// @Produce json
// @Success 200 {object} pb.AistorageResponse "AI storage retrieved successfully"
// @Failure 400 {string} string "Bad request"
// @Router /learning/ai/storage [get]
func (h *Handler) GetIAstorage(ctx *gin.Context) {
	var req pb.AistorageRequest
	jwtToken := ctx.Request.Header.Get("Authourization")
	claims, err := token.ExtractClaim(&config.Config{TokenKey: config.Load().TokenKey}, jwtToken)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	req.Userid = claims["username"].(string)
	res, err := h.LearningService.GetIAstorage(ctx, &req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	ctx.JSON(200, res)
}
