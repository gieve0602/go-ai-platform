package llm

import "context"

type LLMProvider interface {
	Chat(ctx context.Context, req *ChatRequest) (ChatResponse, error)
}
