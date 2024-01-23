package model

type User struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;type:varchar(64);not null;default:'';comment:'姓名'"`
	Age  int64  `json:"age" gorm:"column:age;type:int(64);not null;default:0;comment:'年龄'"`
}

func (e *User) TableName() string {
	return "user"
}
