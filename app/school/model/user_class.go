package model

type UserClass struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;type:varchar(64);not null;default:'';comment:'班级名称'"`
	Room string `json:"room" gorm:"column:room;type:varchar(64);not null;default:'';comment:'班级位置'"`
}

func (e *UserClass) TableName() string {
	return "user_class"
}
