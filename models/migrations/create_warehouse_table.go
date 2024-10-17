package migrations

import "gorm.io/gorm"

type Warehouse struct {
	WhsPK   int    `gorm:"primaryKey"`
	WhsName string `gorm:"type:varchar(255);not null" json:"warehouse_name"`
}

func CreateWarehouseTable(db *gorm.DB) error {
	return db.AutoMigrate(&Warehouse{})
}

func DropWarehouseTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&Warehouse{})
}
