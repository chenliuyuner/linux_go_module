package main

import (
	"gorm_mybatis/config"
	"gorm_mybatis/model"
	"gorm_mybatis/repository"
	"log"
)

func main() {
	db1:=config.InitDb()
	db := repository.BaseRepository[model.User]{DB:db1}
	user := model.User{Name: "Alice_test", Email: "alice_test@example.com"}
	if err := db.Create(&user); err != nil {
		log.Printf("❌ 插入单个数据失败: %v", err)
	} else {
		log.Printf("✅ 插入单个数据成功: %+v", user)
	}

	// 批量插入
	value := []*model.User{
		{Name: "yuorm", Email: "yuorm@example.com"},
		{Name: "peter", Email: "peter@example.com"},
		{Name: "hou", Email: "hou@example.com"},
		{Name: "qin", Email: "qin@example.com"},
		{Name: "Alice", Email: "qin2@example.com"},
	}
	if err := db.Creates(value, 5); err != nil {
		log.Printf("❌ 批量插入失败: %v", err)
	} else {
		log.Println("✅ 批量插入成功")
		for _, u := range value {
			log.Printf("🧍 插入用户: %+v", u)
		}
	}

	// 查询 ID = 5
	u, err := db.SelectById(5)
	if err != nil {
		log.Printf("❌ 查询 ID=5 失败: %v", err)
	} else {
		log.Printf("🔍 查询到 ID=5 的用户: %+v", u)
	}

	// 查询全部用户
	users, err := db.SelectAll()
	if err != nil {
		log.Printf("❌ 查询全部失败: %v", err)
	} else {
		log.Println("📋 当前用户列表:")
		for _, user := range users {
			log.Printf("🧍 %+v", user)
		}
	}

	// 修改 ID=5 的用户姓名
	u.Name = "test"
	if _, err := db.UpdateById(&u); err != nil {
		log.Printf("❌ 更新 ID=5 用户失败: %v", err)
	} else {
		log.Printf("✏️ 成功更新 ID=5 用户为: %+v", u)
	}

	// 再次查询全部用户
	users, err = db.SelectAll()
	if err != nil {
		log.Printf("❌ 查询全部失败: %v", err)
	} else {
		log.Println("📋 更新后用户列表:")
		for _, user := range users {
			log.Printf("🧍 %+v", user)
		}
	}

	// 删除 ID=5 的用户
	deleted, err := db.DeleteById("116")
	if err != nil {
		log.Printf("❌ 删除 ID=116 失败: %v", err)
	} else {
		log.Printf("🗑️ 成功删除用户: %+v", deleted)
	}

	// 最后一次查询全部用户
	users, err = db.SelectAll()
	if err != nil {
		log.Printf("❌ 查询全部失败: %v", err)
	} else {
		log.Println("📋 删除后用户列表:")
		for _, user := range users {
			log.Printf("🧍 %+v", user)
		}
	}
}