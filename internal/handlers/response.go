package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Error string `json:"error"`
}

type statusResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, errMsg string) {
	log.Printf("[ERROR] %s", errMsg)
	c.JSON(statusCode, errorResponse{errMsg})
}

func newStatusResponse(c *gin.Context, statusCode int, message string) {
	log.Printf("[INFO] %s", message)
	c.JSON(statusCode, statusResponse{message})
}
