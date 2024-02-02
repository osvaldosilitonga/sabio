package main

import (
	"fmt"
	"log"
	"ms-order/configs"
	"ms-order/middlewares"
	"ms-order/repositories"
	service "ms-order/services"
	"net"
	"os"

	pb "ms-order/internal/order"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can't load .env file")
	}
}

func main() {
	port := os.Getenv("ORDER_SERVICE_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	db := configs.InitDB()

	orderRepo := repositories.NewOrderRepository(db)

	productService := service.NewProductService()
	customerService := service.NewCustomerService()
	orderService := service.NewOrderService(orderRepo, customerService, productService)

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(middlewares.NewInterceptorLogger()),
		),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterOrderServiceServer(grpcServer, orderService)

	log.Printf("Starting gRPC listener on port : %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
