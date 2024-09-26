package models


type Order struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	UserID    uint    `json:"user_id" gorm:"references users(id)"`
	ProductID uint    `json:"product_id" gorm:"references products(id)"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	
}

func (Order) TableName() string {
	return "orders"
}

type OrderItem struct {
	UserID    uint    `json:"user_id" gorm:"references users(id)"`
	ProductID uint    `json:"product_id" gorm:"references products(id)"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	
}
