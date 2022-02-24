package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CountResponse contains the number of messages contained in the repository
type CountResponse struct {
	Count int `json:"count"`
}

// Count handles processing of a http response with the current number of messages in the repository
func (p Api) Count(w http.ResponseWriter, _ *http.Request) {
	count := p.repository.Count()
	err := json.NewEncoder(w).Encode(CountResponse{
		Count: count,
	})

	if err != nil {
		fmt.Printf("[GoLang] error encoding count response: %v\n", err)
	}
}
