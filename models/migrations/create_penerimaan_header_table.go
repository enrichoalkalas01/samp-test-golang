package migrations

import (
	"time"

	"gorm.io/gorm"
)

type PenerimaanBarangHeader struct {
	TrxInPK      int    `gorm:"primaryKey"`
	TrxInNo      string `gorm:"type:varchar(255);not null"`
	WhsIdf       int
	TrxInDate    time.Time `gorm:"type:date"`
	TrxInSuppIdf int
	TrxInNotes   string                   `gorm:"type:varchar(255)"`
	Details      []PenerimaanBarangDetail `gorm:"foreignKey:TrxInIDF;references:TrxInPK"`
}

func CreatePenerimaanBarangHeaderTable(db *gorm.DB) error {
	return db.AutoMigrate(&PenerimaanBarangHeader{})
}

func DropPenerimaanBarangHeaderTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&PenerimaanBarangHeader{})
}
