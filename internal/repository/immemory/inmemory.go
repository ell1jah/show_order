package immemory

import (
	"sync"

	"github.com/ell1jah/show_order/internal/model"
	"github.com/ell1jah/show_order/internal/repository"
	"github.com/mohae/deepcopy"
)

type orderRepository struct {
	orders []model.Order
	mu     sync.RWMutex
}

func NewRepository() repository.Repository {
	return &orderRepository{
		orders: []model.Order{},
		mu:     sync.RWMutex{},
	}
}

func (or *orderRepository) GetByUID(uid string) (*model.Order, error) {
	or.mu.RLock()
	defer or.mu.RUnlock()

	for _, order := range or.orders {
		if order.OrderUID == uid {
			return (deepcopy.Copy(&order)).(*model.Order), nil
		}
	}

	return nil, model.ErrNotFound
}

func (or *orderRepository) GetAll() ([]model.Order, error) {
	or.mu.RLock()
	defer or.mu.RUnlock()

	return (deepcopy.Copy(or.orders)).([]model.Order), nil
}

func (or *orderRepository) Create(order *model.Order) error {
	or.mu.Lock()
	defer or.mu.Unlock()

	or.orders = append(or.orders, (deepcopy.Copy(*order)).(model.Order))
	return nil
}
