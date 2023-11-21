package entity

type OderRepositoryInterface interface {
	Save(order *Order) error
	GetTotalTransactions() (int, error)
}
