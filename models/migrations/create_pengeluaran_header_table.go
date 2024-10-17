package migrations

import (
	"time"

	"gorm.io/gorm"
)

type PengeluaranBarangHeader struct {
	TrxOutPK      int    `gorm:"primaryKey"`
	TrxOutNo      string `gorm:"type:varchar(255);not null"`
	WhsIdf        int
	TrxOutDate    time.Time `gorm:"type:date"`
	TrxOutSuppIdf int
	TrxOutNotes   string `gorm:"type:varchar(255)"`
}

func CreatePengeluaranBarangHeaderTable(db *gorm.DB) error {
	return db.AutoMigrate(&PengeluaranBarangHeader{})
}

func DropPengeluaranBarangHeaderTable(db *gorm.DB) error {
	return db.Migrator().DropTable(&PengeluaranBarangHeader{})
}
