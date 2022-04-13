/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 20:31:31
 */

package open

import (
	"framework/db"
	"service/enum"
)

// Open 入驻
type Open struct {
	db.Model                               // 通用参数
	UserId         string                  `json:"userId"`         // 用户 id
	ApprovalStatus enum.OpenApprovalStatus `json:"approvalStatus"` // 审核状态
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
