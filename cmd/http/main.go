package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderpb "github.com/crisyantoparulian/grpc-order-service/gen/go/proto/order/v1"
)

func main() {
	// Create a context
	ctx := context.Background()
	mux := runtime.NewServeMux()

	// Configure connection to gRPC server
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Register the gateway
	err := orderpb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatal("Failed to register gateway:", err)
	}

	// Start HTTP server
	addr := ":8080"
	log.Printf("HTTP gateway server running on %s", addr)
	log.Printf("Try: curl http://localhost%s/v1/orders/123", addr)
	log.Printf("Try: curl -X POST http://localhost%s/v1/orders -H 'Content-Type: application/json' -d '{\"product_id\": \"prod123\", \"quantity\": 2}'", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
