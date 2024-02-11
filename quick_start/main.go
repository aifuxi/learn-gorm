package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test-db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("连接数据库失败: ", err)
	}

	// 迁移 schema
	//err = db.AutoMigrate(&Product{})
	//if err != nil {
	//	log.Fatalln("迁移 schema 失败: ", err)
	//}

	// 创建
	db.Create(&Product{
		Code:  "NO.996",
		Price: 60,
	})
	var product Product
	db.First(&product, 3)
	log.Printf("%v\n", product)

	// 更新单个字段
	db.Model(&product).Update("Price", 999)
	log.Printf("%v\n", product)

	// 更新多个字段 - struct
	db.Model(&product).Updates(Product{Price: 200, Code: "FF42"})
	log.Printf("%v\n", product)

	// 更新多个字段 - map
	db.Model(&product).Updates(map[string]any{"Price": 888, "Code": "GG Bound 2"})
	log.Printf("%v\n", product)

	// 删除
	//db.Delete(&product, 1)

	db.First(&product, 3)
	log.Printf("%v\n", product)
}
