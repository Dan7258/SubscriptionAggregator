package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"subagg/internal/models"
)

func (h *Handlers) GetSubscriptions(c *gin.Context) {
	subs, err := h.db.GetSubscriptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, subs)
}

func (h *Handlers) GetSubscriptionByID(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.ParseUint(ids, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sub, err := h.db.GetSubscriptionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, sub)
}

func (h *Handlers) UpdateSubscriptionByID(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.ParseUint(ids, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var sub models.Subscription
	err = c.ShouldBindJSON(&sub)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.db.UpdateSubscriptionByID(id, sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "subscription updated"})

}

func (h *Handlers) CreateSubscription(c *gin.Context) {
	var sub models.Subscription
	err := c.ShouldBindJSON(&sub)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.db.CreateSubscription(sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated})
}

func (h *Handlers) DeleteSubscriptionByID(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.ParseUint(ids, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.db.DeleteSubscriptionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "subscription deleted"})
}

func (h *Handlers) GetSubscriptionsByFilters(c *gin.Context) {
	var filters models.SubscriptionFilters
	err := c.ShouldBindJSON(&filters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	subs, err := h.db.GetSubscriptionsByFilters(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, subs)
}
