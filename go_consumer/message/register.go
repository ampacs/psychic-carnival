package message

import "fmt"

// Register represents a listener that registers messages to a repository received through a channel
type Register struct {
	repository *Repository
}

// NewRegister returns a Register that will register messages to the given repository
func NewRegister(r *Repository) Register {
	return Register{
		repository: r,
	}
}

// RegisterMessages handles registering messages to the repository
func (r *Register) RegisterMessages(mc <-chan Message) <-chan struct{} {
	write := make(chan struct{})
	go func() {
		for {
			message := <-mc
			r.repository.RegisterMessage(message)

			fmt.Printf("[GoLang] Message registered: %+v\n", message)
		}
	}()

	return write
}
