package model

type Menu struct {
	Id    int64
	Name  string  `gorm:"column:name"`
	Price float64 `gorm:"column:price"`
	Pid   int64   `gorm:"column:pid"`
	Count int64   `gorm:"column:count"`
}

func (u Menu) TableName() string {
	return "menu"
}
