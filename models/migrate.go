package models

import (
	"fmt"

	"github.com/enrichoalkalas01/samp-test-golang/models/migrations"
)

// Fungsi untuk mengecek dan migrasi tabel jika belum ada
func MigrateDB() {
	db := GetDB()

	// Cek dan migrate tabel Supplier
	if !db.Migrator().HasTable(&migrations.Supplier{}) {
		migrations.CreateSupplierTable(db)
	} else {
		fmt.Println("Table Supplier already exists")
	}

	// Cek dan migrate tabel Customer
	if !db.Migrator().HasTable(&migrations.Customer{}) {
		migrations.CreateCustomerTable(db)
	} else {
		fmt.Println("Table Customer already exists")
	}

	// Cek dan migrate tabel Product
	if !db.Migrator().HasTable(&migrations.Product{}) {
		migrations.CreateProductTable(db)
	} else {
		fmt.Println("Table Product already exists")
	}

	// Cek dan migrate tabel Warehouse
	if !db.Migrator().HasTable(&migrations.Warehouse{}) {
		migrations.CreateWarehouseTable(db)
	} else {
		fmt.Println("Table Warehouse already exists")
	}

	// Cek dan migrate tabel PenerimaanBarangHeader
	if !db.Migrator().HasTable(&migrations.PenerimaanBarangHeader{}) {
		migrations.CreatePenerimaanBarangHeaderTable(db)
	} else {
		fmt.Println("Table PenerimaanBarangHeader already exists")
	}

	// Cek dan migrate tabel PenerimaanBarangDetail
	if !db.Migrator().HasTable(&migrations.PenerimaanBarangDetail{}) {
		migrations.CreatePenerimaanBarangDetailTable(db)
	} else {
		fmt.Println("Table PenerimaanBarangDetail already exists")
	}

	// Cek dan migrate tabel PengeluaranBarangHeader
	if !db.Migrator().HasTable(&migrations.PengeluaranBarangHeader{}) {
		migrations.CreatePengeluaranBarangHeaderTable(db)
	} else {
		fmt.Println("Table PengeluaranBarangHeader already exists")
	}

	// Cek dan migrate tabel PengeluaranBarangDetail
	if !db.Migrator().HasTable(&migrations.PengeluaranBarangDetail{}) {
		migrations.CreatePengeluaranBarangDetailTable(db)
	} else {
		fmt.Println("Table PengeluaranBarangDetail already exists")
	}
}

// Fungsi untuk menghapus semua tabel (rollback)
func DropDB() {
	db := GetDB()

	migrations.DropSupplierTable(db)
	migrations.DropCustomerTable(db)
	migrations.DropProductTable(db)
	migrations.DropWarehouseTable(db)
	migrations.DropPenerimaanBarangHeaderTable(db)
	migrations.DropPenerimaanBarangDetailTable(db)
	migrations.DropPengeluaranBarangHeaderTable(db)
	migrations.DropPengeluaranBarangDetailTable(db)
}
