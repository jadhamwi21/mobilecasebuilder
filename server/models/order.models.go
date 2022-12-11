package models

type OrderStatus int

const (
	OrderPending OrderStatus = iota
	OrderInProgress
	OrderDone
)

func (orderStatus OrderStatus) String() string {
	return []string{"pending", "in_progress", "packed"}[orderStatus]
}
