package model

type UserRole struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;type:varchar(64);not null;default:'';comment:'角色名称'"`
}

func (e *UserRole) TableName() string {
	return "user_role"
}
