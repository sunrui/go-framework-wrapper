/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 17:24:22
 */

package mysql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// AutoMigrate 创建表
func AutoMigrate(dst ...any) {
	if err := Inst.DB.AutoMigrate(dst...); err != nil {
		panic(err.Error())
	}
}

// Save 插入
func Save(value any) {
	if tx := Inst.DB.Save(value); tx.Error != nil {
		panic(tx.Error.Error())
	}
}

// Truncate 清空数据
func Truncate(dst any) {
	Inst.DB.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&dst)
}

// FindById 根据 id 查找
func FindById[T any](id string) *T {
	var dst T

	if db := Inst.DB.Where("id = ?", id).Find(&dst); db.Error != nil {
		panic(db.Error.Error())
	} else if db.RowsAffected == 1 {
		return &dst
	} else {
		return nil
	}
}

// FindOne 根据条件查找一个
func FindOne[T any](query interface{}, args ...interface{}) *T {
	var dst []T

	if db := Inst.DB.Limit(2).Where(query, args).Find(&dst); db.Error != nil {
		panic(db.Error.Error())
	} else if db.RowsAffected > 1 {
		panic(errors.New(fmt.Sprintf("find %d record", db.RowsAffected)))
	} else if db.RowsAffected == 1 {
		return &dst[0]
	} else {
		return nil
	}
}

// FindPage 根据条件查找分页一个或多个
func FindPage[T any](page int, pageSize int, order string, query interface{}, args ...interface{}) []T {
	var dst []T

	var db *gorm.DB

	if query != nil {
		db = Inst.DB.Order(order).Offset(page*pageSize).Limit(pageSize).Offset(page*pageSize).Where(query, args).Find(&dst)
	} else {
		db = Inst.DB.Order(order).Offset(page * pageSize).Limit(pageSize).Offset(page * pageSize).Find(&dst)
	}

	if db.Error != nil {
		panic(db.Error.Error())
	} else {
		return dst
	}
}

// SoftDeleteById 根据 id 删除
func SoftDeleteById[T any](id string) bool {
	var dst T

	if r := Inst.DB.Where("id = ?", id).Delete(&dst); r.Error != nil {
		panic(r.Error.Error())
	} else if r.RowsAffected >= 1 {
		return true
	} else {
		return false
	}
}

// DeleteById 根据 id 删除
func DeleteById[T any](id string) bool {
	var dst T

	if r := Inst.DB.Unscoped().Where("id = ?", id).Delete(&dst); r.Error != nil {
		panic(r.Error.Error())
	} else if r.RowsAffected >= 1 {
		return true
	} else {
		return false
	}
}
