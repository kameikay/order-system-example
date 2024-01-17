package usecase

import (
	"github.com/kameikay/order-system-example/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

func NewListOrdersUseCase(
	orderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return []ListOrderOutputDTO{}, err
	}

	var ordersOutput []ListOrderOutputDTO
	for _, order := range orders {
		ordersOutput = append(ordersOutput, ListOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ordersOutput, nil
}
