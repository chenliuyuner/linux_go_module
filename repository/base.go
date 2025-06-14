package repository

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/gorm"
)
// Go 导入包时不能访问小写开头的类型
type BaseRepository[T any] struct {
	DB *gorm.DB
}

// [T any] 是 Go 泛型函数
// 插入数据
func (db *BaseRepository[T]) Create(obj *T) error {
	return db.DB.Create(&obj).Error
}

// 批量插入数据
func (db *BaseRepository[T]) Creates(obj []*T, batches int) error {
	return db.DB.CreateInBatches(&obj, batches).Error
}

// 根据 ID 查询
// 返回值
func (db *BaseRepository[T]) SelectById(Id int) (T, error) {
	var obj T
	err := db.DB.Find(&obj, Id).Error
	return obj, err
}

// 返回地址
//
//	func (db *baseRepository[T]) selectById(Id int)(T,error){
//		var obj T
//		err:= db.DB.Find(&obj,Id).Error
//		return obj,err
//	}
//
// 根据批量 ID 查询
// 返回值
func (db *BaseRepository[T]) SelectByIds(idStrs ...string) ([]T, error) {
	//支持多个字符串 ID（自动转 int）
	var ids []int
	for _, s := range idStrs {
		id, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("无效的 ID: %v", s)
		}
		ids = append(ids, id)
	}
	var args []interface{}
	for _, id := range ids {
		args = append(args, id)
	}
	var obj []T
	err := db.DB.Find(&obj, args...).Error
	return obj, err
}

//	func (db *baseRepository[T]) selectByIds(ids []string)([]T,error){
//		var args []interface{}
//		for _,id:= range ids{
//			args = append(args, id)
//		}
//		var obj []T
//		err:= db.DB.Find(&obj,args...).Error
//		return obj,err
//	}
//
// 返回地址
//
//	func (db *baseRepository[T]) selectByIds([]string)([]*T,error){
//		var args []interface{}
//		var ids []string
//		for _,id:= range ids{
//			args = append(args, id)
//		}
//		var obj []*T
//		err:= db.DB.Find(&obj,args...).Error
//		return obj,err
//	}
//
// 查询全部
func (db *BaseRepository[T]) SelectAll() ([]T, error) {
	var objs []T
	err := db.DB.Find(&objs).Error
	return objs, err
}

// 删除
func (db *BaseRepository[T]) DeleteById(id string) (T, error) {
	var obj T
	err := db.DB.Find(&obj, id).Delete(&obj).Error
	return obj, err
}
func (db *BaseRepository[T]) DeleteByIds(idStrs ...string) ([]T, error) {
	//支持多个字符串 ID（自动转 int）
	var ids []int
	for _, s := range idStrs {
		id, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("无效的 ID: %v", s)
		}
		ids = append(ids, id)
	}
	var args []interface{}
	for _, id := range ids {
		args = append(args, id)
	}
	var obj []T
	err := db.DB.Find(&obj, args...).Delete(&obj).Error
	return obj, err
}

// 更新（根据主键）
func (r *BaseRepository[T]) UpdateById(obj *T) (*T, error) {
	err := r.DB.Save(obj).Error
	return obj, err
}
func (r *BaseRepository[T]) UpdateByIds(objs []*T) ([]*T, error) {
	var errList []error
	for _, obj := range objs {
		if err := r.DB.Save(obj).Error; err != nil {
			errList = append(errList, err)
		}
	}
	if len(errList) > 0 {
		return objs, fmt.Errorf("update finished with %d errors", len(errList))
	}
	return objs, nil
}
