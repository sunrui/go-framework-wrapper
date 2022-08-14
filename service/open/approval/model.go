/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 20:31:31
 */

package approval

import (
	"framework/db"
	"service/open"
)

// OpenApproval 入驻审核
type OpenApproval struct {
	db.Model                           // 通用参数
	OpenId         string              `json:"openId"` // 入驻 id
	Open           open.Open           `json:"open" gorm:"foreignKey:OpenId"` // 入驻引用
	UserId         string              `json:"userId"` // 用户 id
	ApprovalStatus open.ApprovalStatus `json:"approvalStatus"` // 审核状态
	Reason         string              `json:"reason"` // 原因
}

// 初始化
func init() {
	// 创建表入驻审核
	if err := db.Mysql.AutoMigrate(&OpenApproval{}); err != nil {
		panic(err.Error())
	}
}
