/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-05 03:20:12
 */

package mysql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// Repository 操作对象
type Repository[T any] struct {
	Mysql *Mysql
}

// NewRepository 创建操作对象
func NewRepository[T any](mysql *Mysql) Repository[T] {
	return Repository[T]{
		Mysql: mysql,
	}
}

// Count 总数
func (repository Repository[T]) Count() (count int64) {
	var dst T
	if db := repository.Mysql.Model(dst).Count(&count); db.Error != nil {
		panic(db.Error.Error())
	}

	return
}

// FindById 根据 id 查找
func (repository Repository[T]) FindById(id string) *T {
	var dst T

	if db := repository.Mysql.DB.Where("id = ?", id).Find(&dst); db.Error != nil {
		panic(db.Error.Error())
	} else if db.RowsAffected == 1 {
		return &dst
	} else {
		return nil
	}
}

// FindOne 根据条件查找一个
func (repository Repository[T]) FindOne(query interface{}, args ...interface{}) *T {
	var dst []T

	if query == nil {
		panic(errors.New("query param cannot be nil"))
	}

	if db := repository.Mysql.DB.Limit(2).Where(query, args...).Find(&dst); db.Error != nil {
		panic(db.Error.Error())
	} else if db.RowsAffected > 1 {
		panic(errors.New(fmt.Sprintf("find %d record", db.RowsAffected)))
	} else if db.RowsAffected == 1 {
		return &dst[0]
	} else {
		return nil
	}
}

// Page 分页
type Page struct {
	Page     int `json:"page" form:"page" validate:"required,gte=1,lte=9999"`        // 分页，从 1 开始
	PageSize int `json:"pageSize" form:"pageSize" validate:"required,gte=1,lte=100"` // 分页大小，最大 100
	//Level      *glog.Level `json:"level" form:"level" validate:"omitempty,oneof=Debug Info Warn Error"` // 日志级别
}

// FindPage 根据条件查找分页一个或多个
func (repository Repository[T]) FindPage(page Page, order string, query interface{}, args ...interface{}) []T {
	var dst []T

	var db *gorm.DB

	db = repository.Mysql.DB.Order(order).Offset((page.Page - 1) * page.PageSize).Limit(page.PageSize)

	if query != nil {
		db = db.Where(query, args...)
	}

	db = db.Find(&dst)

	if db.Error != nil {
		panic(db.Error.Error())
	} else {
		return dst
	}
}

// SoftDeleteById 根据 id 删除
func (repository Repository[T]) SoftDeleteById(id string) bool {
	var dst T

	if r := repository.Mysql.DB.Where("id = ?", id).Delete(&dst); r.Error != nil {
		panic(r.Error.Error())
	} else if r.RowsAffected >= 1 {
		return true
	} else {
		return false
	}
}

// DeleteById 根据 id 删除
func (repository Repository[T]) DeleteById(id string) bool {
	var dst T

	if r := repository.Mysql.DB.Unscoped().Where("id = ?", id).Delete(&dst); r.Error != nil {
		panic(r.Error.Error())
	} else if r.RowsAffected >= 1 {
		return true
	} else {
		return false
	}
}
