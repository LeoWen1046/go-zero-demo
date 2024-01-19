package model

type Office struct {
	BaseModel
	Name    string `json:"name" gorm:"column:name;type:varchar(64);not null;default:'';comment:'办公室名称'"`
	Address string `json:"address" gorm:"column:address;type:varchar(255);not null;default:'';comment:'办公室地址'"`
}

func (o *Office) TableName() string {
	return "office"
}
