package migrations

import "gorm.io/gorm"

type PenerimaanBarangDetail struct {
	TrxInDPK         int `gorm:"primaryKey"`
	TrxInIDF         int `gorm:"index"`
	TrxInDProductIdf int
	TrxInDQtyDus     int
	TrxInDQtyPcs     int
}

func CreatePenerimaanBarangDetailTable(db *gorm.DB) error {
	return db.AutoMigrate(&PenerimaanBarangDetail{})
}

func DropPenerimaanBarangDetailTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&PenerimaanBarangDetail{})
}
