package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"subagg/internal/models"
)

// GetSubscriptions godoc
// @Summary      Get all subscriptions
// @Description  Retrieve a list of all subscriptions
// @Tags         Subscriptions
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.Subscription
// @Failure      500  {object} errorResponse
// @Router       / [get]
func (h *Handlers) GetSubscriptions(c *gin.Context) {
	subs, err := h.db.GetSubscriptions()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, subs)
}

// GetSubscriptionByID godoc
// @Summary      Get a subscription by ID
// @Description  Retrieve a subscription using its ID
// @Tags         Subscriptions
// @Accept       json
// @Produce      json
// @Param        id   path      uint64  true  "Subscription ID"
// @Success      200  {object}  models.Subscription
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Router       /subscriptions/{id} [get]
func (h *Handlers) GetSubscriptionByID(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.ParseUint(ids, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	sub, err := h.db.GetSubscriptionByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(200, sub)
}

// UpdateSubscriptionByID godoc
// @Summary      Update a subscription by ID
// @Description  Update an existing subscription with the provided ID
// @Tags         Subscriptions
// @Accept       json
// @Produce      json
// @Param        id   path      uint64  true  "Subscription ID"
// @Param        subscription  body      models.Subscription  true  "Subscription data"
// @Success      200  {object}  statusResponse
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /subscriptions/{id} [patch]
func (h *Handlers) UpdateSubscriptionByID(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.ParseUint(ids, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var sub models.Subscription
	err = c.ShouldBindJSON(&sub)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.db.UpdateSubscriptionByID(id, sub)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newStatusResponse(c, http.StatusOK, "subscription updated")

}

// CreateSubscription godoc
// @Summary      Create a new subscription
// @Description  Create a new subscription with the provided data. All fields (service_name, price, user_id, start_date) are required.
// @Tags         Subscriptions
// @Accept       json
// @Produce      json
// @Param        subscription  body      models.Subscription  true  "Subscription data"
// @Success      201  {object}  statusResponse
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /subscriptions [post]
func (h *Handlers) CreateSubscription(c *gin.Context) {
	var sub models.Subscription
	err := c.ShouldBindJSON(&sub)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if sub.ServiceName == "" {
		newErrorResponse(c, http.StatusBadRequest, "service_name is required")
		return
	}
	if sub.Price <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "price is required and must be positive")
		return
	}
	if sub.UserID == uuid.Nil {
		newErrorResponse(c, http.StatusBadRequest, "user_id is required")
		return
	}
	if sub.StartDate.String() == "" {
		newErrorResponse(c, http.StatusBadRequest, "start_date is required and must be in MM-YYYY format")
		return
	}
	err = h.db.CreateSubscription(sub)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newStatusResponse(c, http.StatusCreated, "subscription created")
}

// DeleteSubscriptionByID godoc
// @Summary      Delete a subscription by ID
// @Description  Delete a subscription using its ID
// @Tags         Subscriptions
// @Accept       json
// @Produce      json
// @Param        id   path      uint64  true  "Subscription ID"
// @Success      200  {object}  statusResponse
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /subscriptions/{id} [delete]
func (h *Handlers) DeleteSubscriptionByID(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.ParseUint(ids, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.db.DeleteSubscriptionByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newStatusResponse(c, http.StatusOK, "subscription deleted")
}

// GetSubscriptionsByFilters godoc
// @Summary      Get subscriptions by filters
// @Description  Retrieve subscriptions based on provided filters
// @Tags         Subscriptions
// @Accept       json
// @Produce      json
// @Param        filters  body      models.SubscriptionFilters  true  "Subscription filters"
// @Success      200  {array}   models.Subscription
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /subscriptions/filters [post]
func (h *Handlers) GetSubscriptionsByFilters(c *gin.Context) {
	var filters models.SubscriptionFilters
	err := c.ShouldBindJSON(&filters)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	subs, err := h.db.GetSubscriptionsByFilters(filters)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, subs)
}
