package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	orderpb "github.com/crisyantoparulian/grpc-order-service/gen/go/proto/order/v1"

	"github.com/crisyantoparulian/grpc-order-service/internal/auth"
	"github.com/crisyantoparulian/grpc-order-service/internal/interceptor"
	"github.com/crisyantoparulian/grpc-order-service/internal/order"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			interceptor.UnaryAuthInterceptor(auth.ValidateJWT),
		),
	)
	orderpb.RegisterOrderServiceServer(grpcServer, order.NewServer())

	log.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
