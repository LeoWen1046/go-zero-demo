package model

import "time"

// BaseModel 数据库表默认的五个字段
type BaseModel struct {
	ID        int64     `json:"id" gorm:"column:id;type:bigint(20);auto_increment;not null;primary_key;comment:'主键'"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;type:datetime(0);autoCreateTime;comment:'创建时间'"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime(0);autoUpdateTime;comment:'更新时间'"`
	Creator   string    `json:"creator" gorm:"column:creator;type:varchar(64);not null;default:'';comment:'创建人'"`
	Updater   string    `json:"updater" gorm:"column:updater;type:varchar(64);not null;default:'';comment:'更新人'"`
	Version   time.Time `json:"version" gorm:"column:version;type:datetime(0);autoCreateTime;comment:'版本号'"`
}
