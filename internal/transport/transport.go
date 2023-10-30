package transport

import (
	"clientAppService/internal/postgresql/models"
	"clientAppService/internal/service"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type CreateOrderRequest struct {
	AuthorID int `json:"Author ID"`

	FirstAddress  string `json:"First Address"`
	SecondAddress string `json:"Second Address"`
	DeliveryItem  string `json:"Delivery Item"`
}

type CreateOrderResponse struct {
	Order models.Order
}

func MakeCreateOrderEndpoint(cas service.ClientAppService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateOrderRequest)
		createdOrder := cas.Create(ctx, req.AuthorID, req.FirstAddress, req.SecondAddress, req.DeliveryItem)
		return CreateOrderResponse{Order: *createdOrder}, nil
	}
}

func DecodeCutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
