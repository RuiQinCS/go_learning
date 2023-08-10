package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	//gorm.Model
	ID    uint64 `gorm:"primaryKey,autoIncrement"`
	Code  string `gorm:"unique"`
	Price uint
}

func main() {
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open("root:your_password@tcp(localhost:3306)/webook"))
	if err != nil {
		panic("failed to connect database")
	}

	db = db.Debug() // 打印sql日志

	// 迁移 schema
	// 建表
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}

	// Create exp 1
	db.Create(&Product{Code: "D42", Price: 100})

	// Read exp 1
	{
		var product Product
		db.First(&product, 1) // 根据整型主键查找
		fmt.Println("read exp 1. ", product)
	}

	// Update - exp 1 : 将 product 的 price 更新为 200
	{
		var product Product
		db.Model(&product).Where("id = ?", 1).Update("Price", 200)
	}

	{
		var product Product
		db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
		fmt.Println("read exp 2. ", product)
	}

	// Update - exp 2 : 更新多个字段
	{
		var product Product
		product.ID = 1
		// 会更新 Price 和 Code 两个字段  SET `price`=200, `code`='F42'
		db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段

	}

	// Read exp 2
	{
		var product Product
		product.ID = 1
		db.Model(&product).Updates(map[string]interface{}{"Price": 300, "Code": "F42"})
	}

	// Delete - exp 1 : 删除 product
	{
		var product Product
		db.Delete(&product, 1)
	}
}
