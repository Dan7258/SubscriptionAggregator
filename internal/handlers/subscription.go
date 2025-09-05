package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"subagg/internal/models"
)

func (h *Handlers) GetSubscriptions(c *gin.Context) {
	c.JSON(200, h.db.GetSubscriptions)
}

func (h *Handlers) CreateSubscription(c *gin.Context) {
	var sub models.Subscription
	err := c.ShouldBindJSON(&sub)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = h.db.CreateSubscription(sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"subscription": sub})
}
