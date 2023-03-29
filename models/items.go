package models

import "time"

type Items struct {
	ID          uint      `json:"item_id" gorm:"primaryKey;autoIncrement;not null"`
	Item_code   string    `json:"item_code" gorm:"type:varchar(100);unique,not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Quantity    int       `json:"quantity" gorm:"not null"`
	Order_id    uint      `json:"order_id" gorm:"not null;foreignKey:order_id"`
	Created_at  time.Time `json:"created_at" gorm:"default:current_timestamp;not null"`
	Updated_at  time.Time `json:"updated_at" gorm:"null"`
}
