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

	FisrtPhone   string         `json:"First Phone"`
	FirstAddress models.Address `json:"First Address"`

	SecondPhone   string         `json:"Second Phone"`
	SecondAddress models.Address `json:"Second Address"`

	CategoryItem string `json:"Category Item"`
	ItemWeight   string `json:"Item Weight"`
}

type CreateOrderResponse struct {
	Order models.Order
}

func MakeCreateOrderEndpoint(cas service.ClientAppService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateOrderRequest)
		createdOrder := cas.CreateOrder(ctx, req.AuthorID, req.FirstAddress, req.SecondAddress,
			req.CategoryItem, req.ItemWeight, req.FisrtPhone, req.SecondPhone)
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
