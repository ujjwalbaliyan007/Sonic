package services

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type GeminiService struct {
	client *genai.Client
}

func NewGeminiService(apiKey string) (*GeminiService, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}
	return &GeminiService{client: client}, nil
}

func (s *GeminiService) GenerateContent(ctx context.Context, prompt string) (string, error) {
	result, err := s.client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}
	return result.Text(), nil
}
