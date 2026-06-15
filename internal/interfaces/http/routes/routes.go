package routes

import (
	"github.com/gin-gonic/gin"
	"real-estate-agency-onion/internal/application/interfaces"
	"real-estate-agency-onion/internal/interfaces/http/handlers"
	"real-estate-agency-onion/internal/interfaces/http/middleware"
)

func Setup(
	agentH *handlers.AgentHandler,
	tm interfaces.TokenManager,
) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	api.Use(middleware.AuthMiddleware(tm))
	{
		agents := api.Group("/agents")
		{
			agents.POST("", agentH.Create)
			agents.GET("", agentH.List)
			agents.GET("/:id", agentH.GetByID)
		}
	}
	return r
}