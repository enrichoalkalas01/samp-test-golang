package migrations

import "gorm.io/gorm"

type Customer struct {
	CustomerPK   int    `gorm:"primaryKey"`
	CustomerName string `gorm:"type:varchar(255);not null" json:"customer_name"`
}

func CreateCustomerTable(db *gorm.DB) error {
	return db.AutoMigrate(&Customer{})
}

func DropCustomerTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&Customer{})
}
