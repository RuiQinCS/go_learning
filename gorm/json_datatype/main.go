package main

import (
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	//gorm.Model
	ID          uint64 `gorm:"primaryKey,autoIncrement"`
	Code        string `gorm:"unique"`
	Description datatypes.JSON
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

	err = db.Create(&Product{Code: "D47", Description: datatypes.JSON([]byte(`{"name": "product1", "price": 50, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`))}).Error
	if err != nil {
		panic(err)
	}

	{
		var product Product
		db.First(&product, 7) // 根据整型主键查找
		fmt.Println("read exp 1. ", product)
		fmt.Println(product.Description.String())
	}
}
