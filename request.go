package models

type GenerateRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

type GenerateResponse struct {
	Response string `json:"response"`
	Success  bool   `json:"success"`
}

