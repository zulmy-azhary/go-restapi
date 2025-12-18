package handlers

// Handlers contains all handler instances
type Handlers struct {
	Auth    *AuthHandler
	Product *ProductHandler
	Health  *HealthHandler
}

// NewHandlers creates a new Handlers instance
func NewHandlers(auth *AuthHandler, product *ProductHandler, health *HealthHandler) *Handlers {
	return &Handlers{
		Auth:    auth,
		Product: product,
		Health:  health,
	}
}
