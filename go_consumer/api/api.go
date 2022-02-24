package api

// Repository represents the repository contract with which the Api handler will interact
type Repository interface {
	Count() int
	Weights() []int
}

// Api handles processing requests and returning the appropriate responses for API requests
type Api struct {
	repository Repository
}

// NewApi returns a new Api with access to the given repository
func NewApi(r Repository) Api {
	return Api{
		repository: r,
	}
}
