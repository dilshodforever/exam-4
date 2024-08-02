package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"github.com/dilshodforever/4-oyimtixon-api-gatway/api/handler"
	_ "github.com/dilshodforever/4-oyimtixon-api-gatway/docs"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title api gat way
// @version 1.0
// @description Auth service API documentation
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// Middleware setup if needed
	ca, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic(err)
	}

	err = ca.LoadPolicy()
	if err != nil {
		panic(err)
	}
	router := r.Group("/")
	//router.Use(middleware.NewAuth(ca))
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))

	// Game endpoints
	g := router.Group("/game")
	{
		g.GET("/get_levels", h.GetLevels)
		g.POST("/start_level", h.StartLevel)
		g.POST("/complete_level", h.CompleteLevel)
		g.GET("/get_challenge", h.GetChallenge)
		g.POST("/submit_challenge", h.SubmitChallenge)

		g.GET("/get_leaderboard", h.GetLeaderboard)
		g.GET("/get_achievements", h.GetAchievements)
	}
	l := router.Group("/learning")
	{
		l.GET("/topics", h.GetTopics)
		l.GET("/topic/:topic_id", h.GetTopic)
		l.POST("/topic/complete", h.CompleteTopic)
		l.POST("/quiz/submit", h.SubmitQuiz)
		l.GET("/resources", h.GetResources)
		l.POST("/resource/complete", h.CompleteResource)
		l.GET("/progress", h.GetProgress)
		l.GET("/recommendations", h.GetRecommendations)
		l.POST("/feedback/submit", h.SubmitFeedback)
		l.GET("/challenges", h.GetChallenges)
		l.POST("/challenge/solution/submit", h.SubmitChallengeSolution)
		l.GET("/ai/storage", h.GetIAstorage)
	}

	return r
}
