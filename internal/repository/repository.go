package repository

import "github.com/ell1jah/show_order/internal/model"

type Repository interface {
	GetAll() ([]model.Order, error)
	Create(order *model.Order) error
}

type Cache interface {
	Load(orders []model.Order) error
	GetByUID(uid string) (*model.Order, error)
	Create(order *model.Order) error
}
