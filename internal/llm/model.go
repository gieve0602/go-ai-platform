// Package llm
package llm

type Message struct {
	Role    string `json:"role"`
	Content string `json:"context"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	Content string `json:"content"`
}
