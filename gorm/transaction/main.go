package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TransactionPartOne struct {
	ID uint64 `gorm:"primaryKey,autoIncrement"`
}

func (t *TransactionPartOne) TableName() string {
	return "transaction_test_one.test_table_one"
}

type TransactionPartTwo struct {
	ID uint64 `gorm:"primaryKey,autoIncrement"`
}

func (t *TransactionPartTwo) TableName() string {
	return "transaction_test_two.test_table_two"
}

func main() {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(mysql.Open("root:your_pass@tcp(localhost:3306)/webook"))
	if err != nil {
		panic("failed to connect database")
	}

	db = db.Debug() // 打印sql日志

	if err = db.Transaction(func(tx *gorm.DB) error {
		if err0 := tx.AutoMigrate(&TransactionPartOne{}); err0 != nil {
			return err0
		}
		if err0 := tx.AutoMigrate(&TransactionPartTwo{}); err0 != nil {
			return err0
		}
		return nil
	}); err != nil {
		panic(err)
	}
}
