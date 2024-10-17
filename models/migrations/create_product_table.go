package migrations

import "gorm.io/gorm"

type Product struct {
	ProductPK   int    `gorm:"primaryKey"`
	ProductName string `gorm:"type:varchar(255);not null" json:"product_name"`
}

func CreateProductTable(db *gorm.DB) error {
	return db.AutoMigrate(&Product{})
}

func DropProductTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&Product{})
}
