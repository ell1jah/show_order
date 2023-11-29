package logic

import (
	"github.com/ell1jah/show_order/internal/model"
	"github.com/ell1jah/show_order/internal/repository"
	"github.com/pkg/errors"
)

type Logic interface {
	Create(order *model.Order) error
	GetByUID(uid string) (*model.Order, error)
}

type orderLogic struct {
	cache repository.Cache
	repo  repository.Repository
}

func NewRepository(cache repository.Cache, repo repository.Repository) Logic {
	return &orderLogic{
		cache: cache,
		repo:  repo,
	}
}

func (ol *orderLogic) Create(order *model.Order) error {
	err := ol.repo.Create(order)
	if err != nil {
		return errors.Wrap(err, "repository error")
	}

	err = ol.cache.Create(order)
	if err != nil {
		return errors.Wrap(err, "cache error")
	}

	return nil
}

func (ol *orderLogic) GetByUID(uid string) (*model.Order, error) {
	order, err := ol.cache.GetByUID(uid)
	if err != nil {
		return nil, errors.Wrap(err, "cache error")
	}

	return order, nil
}
