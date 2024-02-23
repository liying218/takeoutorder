package model

type Customer struct {
	Id       int64
	Name     string  `gorm:"column:name"`
	Phonenum string  `gorm:"column:phonenum"`
	Address  string  `gorm:"column:address"`
	Sex      string  `gorm:"column:sex"`
	Money    float64 `gorm:"column:money"`
}

func (u Customer) TableName() string {
	return "customer"
}
