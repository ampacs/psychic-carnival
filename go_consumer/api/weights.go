package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AverageWeightResponse contains the averaged weight of all messages in the repository
type AverageWeightResponse struct {
	AverageWeight float64 `json:"average_weight"`
}

// AverageWeight handles processing of a http response with the averaged weight of all messages in the repository
func (p Api) AverageWeight(w http.ResponseWriter, _ *http.Request) {
	weights := p.repository.Weights()
	count := len(weights)
	if count == 0 {
		if err := json.NewEncoder(w).Encode(AverageWeightResponse{}); err != nil {
			fmt.Printf("[GoLang] error encoding empty weight response: %v\n", err)
		}
		return
	}

	var averageWeight float64
	for i := range weights {
		averageWeight += float64(weights[i])
	}
	averageWeight /= float64(count)

	err := json.NewEncoder(w).Encode(AverageWeightResponse{
		AverageWeight: averageWeight,
	})
	if err != nil {
		fmt.Printf("[GoLang] error encoding weight response: %v\n", err)
	}
}
