package service

import (
	"h8-assignment-2/helpers"
	"h8-assignment-2/models"
	"h8-assignment-2/repository"
)

type OrderService interface {
	FindAll() ([]*models.Orders, error)
	CreateOrder(order *models.Orders) (*models.Orders, error)
	UpdateOrder(reqOrder *helpers.RequestOrder) (helpers.ResponseOrder, error)
	DeleteOrder(order *models.Orders) error
}

type orderService struct {
	repository repository.OrderRepository
}

func NewOrderService(repository repository.OrderRepository) OrderService {
	return &orderService{
		repository: repository,
	}
}

func (s *orderService) FindAll() ([]*models.Orders, error) {
	orders, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *orderService) CreateOrder(order *models.Orders) (*models.Orders, error) {
	createdOrder, err := s.repository.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return createdOrder, nil
}

func (s *orderService) UpdateOrder(reqOrder *helpers.RequestOrder) (helpers.ResponseOrder, error) {
	updatedOrder, err := s.repository.UpdateOrder(reqOrder)
	if err != nil {
		return helpers.ResponseOrder{}, err
	}
	return updatedOrder, nil
}

func (s *orderService) DeleteOrder(order *models.Orders) error {
	err := s.repository.DeleteOrder(order)
	if err != nil {
		return err
	}
	return nil
}
