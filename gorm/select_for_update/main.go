package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	//gorm.Model
	ID    uint64 `gorm:"primaryKey,autoIncrement"`
	Code  string `gorm:"unique"`
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:your_pass@tcp(localhost:3306)/crud"))
	if err != nil {
		panic("failed to connect database")
	}

	db = db.Debug() // 打印sql日志

	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}

	db.Create(&Product{Code: "D42", Price: 100})

	var maxid uint
	err = db.Model(&Product{}).Clauses(clause.Locking{Strength: "UPDATE"}).Select("max(price) as price").Find(&maxid).Error
	if err != nil {
		panic(err)
	}

	fmt.Print(maxid)
}
