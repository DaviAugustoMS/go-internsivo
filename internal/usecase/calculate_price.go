package usecase

import "github.com/DaviAugustoMS/go-internsivo/internal/entity"

type OrderInput struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutput struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CaculateFinalPrice struct {
	OderRepository entity.OderRepositoryInterface
}

func NewCalculateFinalPrice(orderRepository entity.OderRepositoryInterface) *CaculateFinalPrice {
	return &CaculateFinalPrice{
		OderRepository: orderRepository,
	}
}

func (c *CaculateFinalPrice) Execute(input OrderOutput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFilnalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
