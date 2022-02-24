package message

import "sync"

// Repository represents a collection of messages
type Repository struct {
	mu       sync.Mutex
	messages []Message
}

// RegisterMessage allows adding a message to the repository
func (r *Repository) RegisterMessage(message Message) {
	r.mu.Lock()
	r.messages = append(r.messages, message)
	r.mu.Unlock()
}

// Count returns the number of registered messages
func (r *Repository) Count() int {
	r.mu.Lock()
	count := len(r.messages)
	r.mu.Unlock()

	return count
}

// Weights returns the list of weights of all registered messages
func (r *Repository) Weights() []int {
	r.mu.Lock()
	weights := make([]int, len(r.messages))
	for i, message := range r.messages {
		weights[i] = message.Weight
	}
	r.mu.Unlock()

	return weights
}
