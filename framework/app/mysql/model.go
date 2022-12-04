/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-06 18:10:03
 */

package mysql

import (
	"framework/app/util"
	"gorm.io/gorm"
	"time"
)

// Model 数据库
type Model struct {
	Id        string         `json:"id" gorm:"primaryKey;type:char(12);comment:主键 id"`    // 主键 id
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime:milli;comment:创建时间"` // 创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime:milli;comment:更新时间"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"comment:删除时间"`            // 删除时间
}

// BeforeCreate 创建前回调
func (model *Model) BeforeCreate(*gorm.DB) (err error) {
	model.Id = util.CreateNanoid(12)
	return nil
}
