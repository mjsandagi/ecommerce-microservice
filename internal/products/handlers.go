package products

import "net/http"
import "github.com/mjsandagi/go-ecommerce/internal/json"
// import "log"

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
	// err := h.service.ListProducts(r.Context())
	// if err != nil {
	// 	log.Println("Error fetching products:", err)
	// 	http.Error(w, error.Error(err), http.StatusInternalServerError) // The server encountered an unexpected error, preventing it from fulfilling the request.
	// }
	// 2. Return JSON response
	products := []string{"Product 1", "Product 2", "Product 3"}
	json.Write(w, http.StatusOK, products)
}