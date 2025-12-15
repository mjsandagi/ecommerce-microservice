package products

import "context"

type Service interface {
	// Add fields as necessary, e.g., a reference to the service layer
	ListProducts(ctx context.Context) error
}

type svc struct {
	// repository
}

func NewService() Service {
	return &svc{
	}
}

func (s *svc) ListProducts(ctx context.Context) (error) {
	return nil
}