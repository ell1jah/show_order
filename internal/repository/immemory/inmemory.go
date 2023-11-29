package immemory

import (
	"sync"

	"github.com/ell1jah/show_order/internal/model"
	"github.com/ell1jah/show_order/internal/repository"
	"github.com/mohae/deepcopy"
)

type orderCache struct {
	orders []model.Order
	mu     sync.RWMutex
}

func NewCache() repository.Cache {
	return &orderCache{
		orders: []model.Order{},
		mu:     sync.RWMutex{},
	}
}

func (oc *orderCache) GetByUID(uid string) (*model.Order, error) {
	oc.mu.RLock()
	defer oc.mu.RUnlock()

	for _, order := range oc.orders {
		if order.OrderUID == uid {
			return (deepcopy.Copy(&order)).(*model.Order), nil
		}
	}

	return nil, model.ErrNotFound
}

func (oc *orderCache) Load(orders []model.Order) error {
	oc.mu.Lock()
	defer oc.mu.Unlock()

	oc.orders = append(oc.orders, (deepcopy.Copy(orders)).([]model.Order)...)
	return nil
}

func (oc *orderCache) Create(order *model.Order) error {
	oc.mu.Lock()
	defer oc.mu.Unlock()

	oc.orders = append(oc.orders, (deepcopy.Copy(*order)).(model.Order))
	return nil
}
