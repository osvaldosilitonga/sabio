package service

import (
	"context"
	"log"
	"time"

	"ms-order/domain/entity"
	"ms-order/helpers"
	pb "ms-order/internal/order"
	"ms-order/repositories"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Order struct {
	product *pb.Order
	pb.UnimplementedOrderServiceServer
	OrderRepo       repositories.Order
	CustomerService Customer
	ProductService  Product
}

func NewOrderService(db repositories.Order, cs Customer, ps Product) *Order {
	return &Order{
		OrderRepo:       db,
		CustomerService: cs,
		ProductService:  ps,
	}
}

func (p *Order) Create(ctx context.Context, req *pb.CreateReq) (*pb.Order, error) {

	order := entity.Order{
		CustomerID: int(req.CustomerId),
		ProductID:  int(req.ProductId),
		Qty:        int(req.Qty),
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// check if costumer_id exists
	_, code, err := p.CustomerService.FindById(ctx, int(req.ProductId))
	if err != nil || code != 200 {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// check if product_id and stock valid
	product, _, err := p.ProductService.FindById(ctx, int(req.ProductId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Println(product, "<------ product")

	order.Total = product.Price * float64(order.Qty)

	// save to database
	insertedId, err := p.OrderRepo.Create(ctx, order)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	order.ID = int(insertedId)
	log.Println(order.ID)

	return helpers.ToOrderResponse(&order), nil
}

func (p *Order) FindById(ctx context.Context, req *pb.FindByIdReq) (*pb.Order, error) {
	return nil, nil
}
