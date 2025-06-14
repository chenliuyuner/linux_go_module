package config

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func InitDb() *gorm.DB{
	DSN:= "root:123@tcp(192.168.103.131)/learn"
	db1,err:= gorm.Open(mysql.Open(DSN),&gorm.Config{})
	if err != nil {
		log.Fatalf("❌ 连接数据库失败: %v", err)
	}
	log.Println("✅ 成功连接数据库")
	return db1
}