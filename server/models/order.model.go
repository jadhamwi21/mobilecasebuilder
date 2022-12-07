package models

type OrderStatusType string

const (
	OrderPending   OrderStatusType = "pending"
	OrderPreparing OrderStatusType = "preparing"
	OrderShipped   OrderStatusType = "shipped"
)

type Order struct {
	Id        string
	Timestamp string
	Status    OrderStatusType
}
