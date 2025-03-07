package handlers

import (
	"net/http"
	"sers/services"

	"github.com/gin-gonic/gin"
)

type SOSHandler struct {
	sosService *services.SOSService
}

func NewSOSHandler(sosService *services.SOSService) *SOSHandler {
	return &SOSHandler{sosService: sosService}
}

func (h *SOSHandler) TriggerSOS(c *gin.Context) {
	var sosRequest struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Message   string  `json:"message"`
	}

	if err := c.BindJSON(&sosRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	err := h.sosService.TriggerEmergency(uint(userID.(float64)), sosRequest.Latitude, sosRequest.Longitude, sosRequest.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SOS triggered successfully"})
}
