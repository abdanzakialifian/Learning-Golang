package test

import "time"

type Product struct {
	Id           string    `gorm:"column:id;primary_key"`
	Name         string    `gorm:"column:name"`
	Price        int64     `gorm:"column:price"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	LikedByUsers []User    `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:product_id;references:id;joinReferences:user_id"`
}

func (product *Product) TableName() string {
	return "product"
}
