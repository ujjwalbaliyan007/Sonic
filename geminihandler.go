package handlers

import (
	"BACKEND/models"
	"BACKEND/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeminiHandler struct {
	service *services.GeminiService
}

func NewGeminiHandler(service *services.GeminiService) *GeminiHandler {
	return &GeminiHandler{service: service}
}

func (h *GeminiHandler) Generate(c *gin.Context) {
	var req models.GenerateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.service.GenerateContent(c.Request.Context(), req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate response"})
		return
	}

	c.JSON(http.StatusOK, models.GenerateResponse{
		Response: response,
		Success:  true,
	})
}

//curl -X POST http://localhost:8080/api/generate -H "Content-Type: application/json" -d '{"prompt": "Tell me a joke in one line."}'
