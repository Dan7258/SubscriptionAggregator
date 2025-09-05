package routes

import (
	"github.com/gin-gonic/gin"
	"subagg/internal/handlers"
)

func RegisterRoutes(r *gin.Engine, h *handlers.Handlers) {
	r.GET("/", h.GetSubscriptions)
	r.GET("/subscriptions/:id", h.GetSubscriptionByID)
	r.POST("/subscriptions", h.CreateSubscription)
	r.POST("/subscriptions/filters", h.GetSubscriptionsByFilters)
	r.PATCH("/subscriptions/:id", h.UpdateSubscriptionByID)
	r.DELETE("/subscriptions/:id", h.DeleteSubscriptionByID)
}
