package models

type Bike struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Model string  `json:"model"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}
