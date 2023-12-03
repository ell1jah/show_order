package postgres

import (
	"fmt"

	"github.com/ell1jah/show_order/internal/model"
	"github.com/ell1jah/show_order/internal/repository"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.Repository {
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) Create(order *model.Order) error {
	err := or.db.Create(&order).Error
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("database error, order:%+v", order))
	}

	return nil
}

func (or *orderRepository) GetAll() ([]model.Order, error) {
	var orders []model.Order
	err := or.db.Preload("Delivery").Preload("Payment").Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, errors.Wrap(err, "database error")
	}

	return orders, err
}
