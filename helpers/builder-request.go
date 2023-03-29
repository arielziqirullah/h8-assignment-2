package helpers

type RequestOrder struct {
	OrderID  uint   `json:"order_id"`
	Customer string `json:"customer_name"`
	Items    []Item `json:"items"`
}

type Item struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
