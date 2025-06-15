package test

import "time"

type User struct {
	Id          string    `gorm:"column:id;primary_key;<-:create"` // "<-:create is field permission"
	Name        string    `gorm:"column:name"`
	Password    string    `gorm:"column:password"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"` // "<-:create is field permission"
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"column:information;-"` // "- is field permission"
}

func (user *User) TableName() string {
	return "user"
}
