/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:26:01
 */

package open

import (
	"framework/db"
)

// Open 入驻
type Open struct {
	db.Model                      // 通用参数
	UserId         string         `json:"userId"`         // 用户 id
	ApprovalStatus ApprovalStatus `json:"approvalStatus"` // 审核状态
}

// 初始化
func init() {
	// 创建表入驻
	if err := db.Mysql.AutoMigrate(&Open{}); err != nil {
		panic(err.Error())
	}
}

// Save 存储
func (open *Open) Save() {
	db.Mysql.Save(open)
}
