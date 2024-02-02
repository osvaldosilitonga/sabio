package helpers

import (
	"ms-order/domain/entity"
	pb "ms-order/internal/order"
)

func ToOrderResponse(data *entity.Order) *pb.Order {
	return &pb.Order{
		Id:         int64(data.ID),
		CustomerId: int64(data.CustomerID),
		ProductId:  int64(data.ProductID),
		Qty:        int64(data.Qty),
		Total:      float32(data.Total),
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}
