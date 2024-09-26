package models


type Product struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	ModelName   string   `json:"name"`
	Description int   `json:"description"`
	Price       float64  `json:"price"`
	CategoryID  int      `json:"category_id" gorm:"references category(id)"`
	// Categories  Category `json:"category" gorm:"foreignKey: CategoryID"`
}

func (Product) TableName() string {
	return "products"
}
