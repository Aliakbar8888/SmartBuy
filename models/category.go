package models


type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"category"  gorm:"unique; not null"`
}

func (Category) TableName() string {
	return "categories"
}
