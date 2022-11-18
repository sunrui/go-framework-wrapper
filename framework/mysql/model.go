/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 18:10:03
 */

package mysql

import (
	"framework/util"
	"gorm.io/gorm"
	"time"
)

// Model 数据库通用对象
type Model[T any] struct {
	Id        string         `json:"id" gorm:"primaryKey;type:char(10);comment:主键 id"`    // 主键 id
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime:milli;comment:创建时间"` // 创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime:milli;comment:更新时间"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"comment:删除时间"`            // 删除时间
}

// BeforeCreate 创建对象前回调
func (model *Model[T]) BeforeCreate(*gorm.DB) (err error) {
	model.Id = util.CreateNanoid(10)
	return nil
}
