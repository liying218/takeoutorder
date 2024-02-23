package model

type Producer struct {
	Id       int64
	Name     string  `gorm:"column:name"`
	Phonenum string  `gorm:"column:phonenum"`
	Address  string  `gorm:"column:address"`
	Money    float64 `gorm:"column:money"`
}

func (u Producer) TableName() string {
	return "producer"
}
