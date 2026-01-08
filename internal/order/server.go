package order

import (
	"context"
	"log"

	orderpb "github.com/crisyantoparulian/grpc-order-service/gen/go/proto/order/v1"
	"github.com/crisyantoparulian/grpc-order-service/internal/interceptor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	userID, ok := interceptor.UserIDFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "user not found in context")
	}

	log.Println("Authenticated user:", userID)

	return &orderpb.GetOrderResponse{
		OrderId:    req.OrderId,
		Status:     "CREATED",
		TotalPrice: 2,
	}, nil
}

func (s *Server) CreateOrder(
	ctx context.Context,
	req *orderpb.CreateOrderRequest,
) (*orderpb.CreateOrderResponse, error) {

	userID, _ := interceptor.UserIDFromContext(ctx)

	log.Printf(
		"CreateOrder user=%s product=%s qty=%d",
		userID,
		req.ProductId,
		req.Quantity,
	)

	return &orderpb.CreateOrderResponse{
		OrderId: "ORDER-1001",
		Status:  "CREATED",
	}, nil
}
