package routes

import (
	"github.com/gin-gonic/gin"
	"subagg/internal/handlers"
)

func RegisterRoutes(r *gin.Engine, h *handlers.Handlers) {
	r.GET("/", h.GetSubscriptions)
}
