package test

type UserLog struct {
	Id        int    `gorm:"column:id;primary_key;autoIncrement"`
	UserId    string `gorm:"column:user_id"`
	Action    string `gorm:"column:action"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (l *UserLog) TableName() string {
	return "user_logs"
}
