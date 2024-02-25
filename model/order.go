package model

type Order struct {
	Id    int64
	Pid   int64   `gorm:"column:pid"`
	Cid   int64   `gorm:"column:cid"`
	Rid   int64   `gorm:"column:rid"`
	Foods string  `gorm:"column:foods"`
	Price float64 `gorm:"column:price"`
	State string  `gorm:"column:state"`
}

func (u Order) TableName() string {
	return "order"
}
