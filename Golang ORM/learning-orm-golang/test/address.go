package test

import "time"

type Address struct {
	Id        int64     `gorm:"column:id;primary_key;autoIncrement"`
	UserId    string    `gorm:"column:user_id"`
	Address   string    `gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User      User      `gorm:"foreignKey:user_id;references:id"`
}

func (address *Address) TableName() string {
	return "address"
}
