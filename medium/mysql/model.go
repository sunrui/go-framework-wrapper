/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 18:10:03
 */

package mysql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"medium/util"
	"time"
)

// Model 数据库通用对象
type Model[T any] struct {
	Id        string         `json:"id" gorm:"primaryKey;type:char(16);comment:主键 id"`    // 主键 id
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime:milli;comment:创建时间"` // 创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime:milli;comment:更新时间"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"comment:删除时间"`            // 删除时间
}

func NewEmptyModel[T any]() Model[T] {
	return Model[T]{}
}

// BeforeCreate 创建对象前回调
func (model *Model[T]) BeforeCreate(*gorm.DB) (err error) {
	model.Id = util.CreateNanoid(16)
	return nil
}

// FindById 根据 id 查找
func (model *Model[T]) FindById(id string) *T {
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
func (model *Model[T]) FindOne(query interface{}, args ...interface{}) *T {
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
func (model *Model[T]) FindPage(page int, pageSize int, order string, query interface{}, args ...interface{}) []T {
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
func (model *Model[T]) SoftDeleteById(id string) bool {
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
func (model *Model[T]) DeleteById(id string) bool {
	var dst T

	if r := Inst.DB.Unscoped().Where("id = ?", id).Delete(&dst); r.Error != nil {
		panic(r.Error.Error())
	} else if r.RowsAffected >= 1 {
		return true
	} else {
		return false
	}
}
