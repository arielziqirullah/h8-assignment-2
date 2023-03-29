package models

import "time"

type Orders struct {
	ID            uint      `json:"order_id" gorm:"primaryKey;autoIncrement;not null"`
	Customer_name string    `json:"customer_name" gorm:"type:varchar(100);not null"`
	Orderd_at     time.Time `json:"orderd_at" gorm:"default:current_timestamp;not null"`
	Created_at    time.Time `json:"created_at" gorm:"default:current_timestamp;not null"`
	Updated_at    time.Time `json:"updated_at" gorm:"null"`
	Items         []Items   `json:"items" gorm:"foreignKey:Order_id"`
}
