package api

type Handlers struct {
	// Add any dependencies your handlers might need, e.g., database connections
}

// You can define methods on Handlers for each of your HTTP endpoints
func NewHandlers() *Handlers {
	return &Handlers{}
}
