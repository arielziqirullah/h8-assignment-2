package repository

import (
	"h8-assignment-2/helpers"
	"h8-assignment-2/models"
	"time"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll() ([]*models.Orders, error)
	CreateOrder(order *models.Orders) (*models.Orders, error)
	UpdateOrder(reqOrder *helpers.RequestOrder) (helpers.ResponseOrder, error)
	DeleteOrder(order *models.Orders) error
}

type orderConnection struct {
	connection *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderConnection{
		connection: db,
	}
}

func (db *orderConnection) FindAll() ([]*models.Orders, error) {
	var orders []*models.Orders
	err := db.connection.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (db *orderConnection) CreateOrder(order *models.Orders) (*models.Orders, error) {
	err := db.connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return order, nil
}

func (db *orderConnection) UpdateOrder(reqOrder *helpers.RequestOrder) (helpers.ResponseOrder, error) {

	var order *models.Orders

	err := db.connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", reqOrder.OrderID).First(&order).Error; err != nil {
			return err
		}

		if err := tx.Model(&order).Updates(&models.Orders{
			Customer_name: reqOrder.Customer,
			Updated_at:    time.Now(),
		}).Error; err != nil {
			return err
		}

		var existingItems []*models.Items
		if err := tx.Where("order_id = ?", reqOrder.OrderID).Find(&existingItems).Error; err != nil {
			return err
		}

		existingItemMap := make(map[string]*models.Items)
		for _, item := range existingItems {
			existingItemMap[item.Item_code] = item
		}

		for _, item := range reqOrder.Items {
			if existingItem, ok := existingItemMap[item.ItemCode]; ok {
				existingItem.Description = item.Description
				existingItem.Quantity = item.Quantity
				existingItem.Updated_at = time.Now()

				if err := tx.Save(existingItem).Error; err != nil {
					return err
				}
			} else {
				newItem := &models.Items{
					Order_id:    reqOrder.OrderID,
					Item_code:   item.ItemCode,
					Description: item.Description,
					Quantity:    item.Quantity,
					Updated_at:  time.Now(),
				}

				if err := tx.Create(newItem).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		return helpers.ResponseOrder{}, err
	}

	err = db.connection.Preload("Items").First(&order, reqOrder.OrderID).Error
	if err != nil {
		return helpers.ResponseOrder{}, err
	}

	orderResponse := helpers.ResponseOrder{
		OrderID:      order.ID,
		CustomerName: order.Customer_name,
		OrderdAt:     order.Orderd_at,
	}

	for _, item := range order.Items {
		orderResponse.Items = append(orderResponse.Items, helpers.ResponseItem{
			ItemCode:    item.Item_code,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}

	return orderResponse, nil
}

func (db *orderConnection) DeleteOrder(order *models.Orders) error {
	err := db.connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("order_id = ?", order.ID).Delete(&models.Items{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(order).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
