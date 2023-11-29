package repository

import "github.com/ell1jah/show_order/internal/model"

type Repository interface {
	GetByUID(uid string) (*model.Order, error)
	GetAll() ([]model.Order, error)
	Create(order *model.Order) error
}
