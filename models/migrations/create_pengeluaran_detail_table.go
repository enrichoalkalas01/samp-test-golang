package migrations

import "gorm.io/gorm"

type PengeluaranBarangDetail struct {
	TrxOutDPK         int `gorm:"primaryKey"`
	TrxOutIDF         int
	TrxOutDProductIdf int
	TrxOutDQtyDus     int
	TrxOutDQtyPcs     int
}

func CreatePengeluaranBarangDetailTable(db *gorm.DB) error {
	return db.AutoMigrate(&PengeluaranBarangDetail{})
}

func DropPengeluaranBarangDetailTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&PengeluaranBarangDetail{})
}
