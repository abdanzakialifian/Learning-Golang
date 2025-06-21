package test

import "time"

type User struct {
	Id          string    `gorm:"column:id;primary_key;<-:create"` // "<-:create is field permission"
	Name        Name      `gorm:"embedded"`
	Password    string    `gorm:"column:password"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"` // "<-:create is field permission"
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Wallet      Wallet    `gorm:"foreignKey:user_id;references:id"`
	Addresses   []Address `gorm:"foreignKey:user_id;references:id"`
	Information string    `gorm:"-"` // "- is field permission"
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}

func (user *User) TableName() string {
	return "user"
}
