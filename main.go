package main

import (
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	//root:abc456@tcp(10.134.123.183:3306)/go_test_db
	db, err := gorm.Open("mysql", "root:abc456@tcp(10.134.123.183:3306)/go_test_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("connect mysql go_test_db error:%v", err)
		return
	}
	log.Println("connect mysql go_test_db success")
	defer db.Close()
}