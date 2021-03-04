package service

import (
	"errors"
	"gb-go-architecture/lesson-2/shop_new/models"
	"gb-go-architecture/lesson-2/shop_new/repository"
	"log"
)

type Service interface {
	CreateOrder(order *models.Order) (*models.Order, error)
}

type Notification interface {
	SendOrderCreated(order *models.Order) error
}

type service struct {
	userNotif  Notification
	staffNotif Notification
	rep        repository.Repository
}

var (
	ErrItemNotExists = errors.New("item not exists")
)

func (s *service) CreateOrder(order *models.Order) (*models.Order, error) {
	for _, itemID := range order.ItemIDs {
		_, err := s.rep.GetItem(itemID)
		if err != nil && err != repository.ErrNotFound {
			return nil, err
		}
		if err == repository.ErrNotFound {
			return nil, ErrItemNotExists
		}
	}

	order, err := s.rep.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	if err := s.staffNotif.SendOrderCreated(order); err != nil {
		log.Printf("send staffNotif error: %s", err)
	}
	if err := s.userNotif.SendOrderCreated(order); err != nil {
		log.Printf("send userNotif error: %s", err)
	}
	return order, err
}

func NewService(rep repository.Repository, staffNotif, userNotif Notification) Service {
	return &service{
		userNotif:  userNotif,
		staffNotif: staffNotif,
		rep:        rep,
	}
}
