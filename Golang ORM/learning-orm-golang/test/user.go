package test

import "time"

type User struct {
	Id        string    `gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:name"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (user *User) TableName() string {
	return "user"
}
