package helpers

import "time"

type ResponseFindAll struct {
	Data interface{} `json:"data"`
}

type ResponseOrder struct {
	OrderID      uint           `json:"order_id"`
	CustomerName string         `json:"customer_name"`
	OrderdAt     time.Time      `json:"orderd_at"`
	Items        []ResponseItem `json:"items"`
}

type ResponseItem struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type ResponseUpdateOrder struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

type ResponseDeleteOrder struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
