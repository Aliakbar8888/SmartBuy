package models


type Wishlist struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	UserID    uint `gorm:"references users(id)" json:"user_id"`
	ProductID uint    `json:"product_id" gorm:"references products(id)"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int  `json:"quantity"`
}
