package test

import "time"

type UserLog struct {
	Id        int       `gorm:"column:id;primary_key;autoIncrement"`
	UserId    string    `gorm:"column:user_id"`
	Action    string    `gorm:"column:action"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (l *UserLog) TableName() string {
	return "user_logs"
}
