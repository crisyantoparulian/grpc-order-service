package order

import (
	"context"

	orderpb "github.com/crisyantoparulian/grpc-order-service/gen/go/order/v1"
)

type Server struct {
	orderpb.UnimplementedOrderServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetOrder(
	ctx context.Context,
	req *orderpb.GetOrderRequest,
) (*orderpb.GetOrderResponse, error) {
	return &orderpb.GetOrderResponse{
		OrderId: req.OrderId,
		Status:  "CREATED",
	}, nil
}
