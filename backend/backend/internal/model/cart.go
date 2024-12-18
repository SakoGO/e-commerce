package model

type Cart1 struct {
	CartItemID      int     `json:"cart_item_id" gorm:"primaryKey"`
	OrderID         int     `json:"order_id" gorm:"foreignKey:OrderID"`
	ProductID       int     `json:"product_id" gorm:"foreignKey:ProductID"`
	Quantity        int     `json:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase"`

	Order   Order   `json:"order"`
	Product Product `json:"product"`
}

// Логика:
// В Cart имеет прямую связь с UserID. Сюда user собирает слайс объектов из сущности Product.
// Cart имеет bool значение is_active, которое отслеживает состояние корзины (наличие в ней product)
// Cart подтягивает из product его price и высчитывает итоговую сумму будущего order
// Так же cart подтягивает значение stock и если оно <= 0, то не позволяет собрать order, т.к корзина пуста
// И т.д......

type Cart struct {
	CartID int `gorm:"primarykey"`
}
