package migrations

import "gorm.io/gorm"

type Supplier struct {
	SupplierPK   int    `gorm:"primaryKey"`
	SupplierName string `gorm:"type:varchar(255);not null" json:"supplier_name"`
}

func CreateSupplierTable(db *gorm.DB) error {
	return db.AutoMigrate(&Supplier{})
}

func DropSupplierTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&Supplier{})
}
