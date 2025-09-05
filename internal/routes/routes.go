package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "subagg/docs"
	"subagg/internal/handlers"
)

func RegisterRoutes(r *gin.Engine, h *handlers.Handlers) {
	r.GET("/", h.GetSubscriptions)
	r.GET("/subscriptions/:id", h.GetSubscriptionByID)
	r.POST("/subscriptions", h.CreateSubscription)
	r.POST("/subscriptions/filters", h.GetSubscriptionsByFilters)
	r.PATCH("/subscriptions/:id", h.UpdateSubscriptionByID)
	r.DELETE("/subscriptions/:id", h.DeleteSubscriptionByID)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
