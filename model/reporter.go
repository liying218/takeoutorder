package model

type Reporter struct {
	Id       int64
	Name     string `gorm:"column:name"`
	Phonenum string `gorm:"column:phonenum"`
	Address  string `gorm:"column:address"`
	Sex      string `gorm:"column:sex"`
	Count    int64  `gorm:"column:count"`
}

func (u Reporter) TableName() string {
	return "reporter"
}
