package entity

type SourceOfFund struct {
	ID   uint64 `gorm:"primaryKey"`
	Name string
}

// func (SourceOfFund) TableName() string {
// 	return "source_of_funds"
// }
