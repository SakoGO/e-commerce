package model

type Cart struct {
	CartItemID      int     `json:"cart_item_id" gorm:"primaryKey"`
	OrderID         int     `json:"order_id" gorm:"foreignKey:OrderID"`
	ProductID       int     `json:"product_id" gorm:"foreignKey:ProductID"`
	Quantity        int     `json:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase"`

	Order   Order   `json:"order"`
	Product Product `json:"product"`
}
