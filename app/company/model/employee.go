package model

type Employee struct {
	BaseModel
	Name     string `json:"name" gorm:"column:name;type:varchar(64);not null;default:'';comment:'员工姓名'"`
	Age      int64  `json:"age" gorm:"column:age;type:int(64);not null;default:0;comment:'员工年龄'"`
	Company  string `json:"company" gorm:"column:company;type:varchar(255);not null;default:'';comment:'公司'"`
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(64);not null;default:'';comment:'电话'"`
	OfficeID int64  `json:"officeId" gorm:"column:office_id;type:int(64);not null;default:0;comment:'办公室id'"`
}

func (e *Employee) TableName() string {
	return "employee"
}
