package repository

import (
	"github.com/jadhamwi21/mobilecasebuilder/domain/entity"
)

type OrderRepository interface {
	SaveOrder(*entity.Order) (map[string]string, error)
	GetOrderById(string) (*entity.Order, error)
	GetAllOrders() ([]*entity.Order, error)
	GetOrderCases(string) ([]*entity.Case, error)
	GetOrdersByUserId(string) ([]*entity.Order, error)
}
