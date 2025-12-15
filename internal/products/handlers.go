package products

import "net/http"
import "encoding/json"

type handler struct {
	service Service
}


func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. Call the service layer --> ListProducts method
	// 2. Return JSON response
	products := []string{"Product 1", "Product 2", "Product 3"}
	json.NewEncoder(w).Encode(products)
}