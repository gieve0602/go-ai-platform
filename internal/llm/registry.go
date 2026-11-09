package llm

import (
	"errors"
	"sync"
)

var ErrProviderNotFound = errors.New("provider not found")

type ProviderRegistry struct {
	providers map[string]LLMProvider
	mux       sync.RWMutex
}

func (pr *ProviderRegistry) Registry(name string, provider LLMProvider) {
	pr.mux.Lock()
	defer pr.mux.Unlock()
	pr.providers[name] = provider
}

func (pr *ProviderRegistry) Get(name string) (LLMProvider, error) {
	pr.mux.RLock()
	defer pr.mux.RUnlock()
	if provider, ok := pr.providers[name]; ok {
		return provider, nil
	}
	return nil, ErrProviderNotFound
}
