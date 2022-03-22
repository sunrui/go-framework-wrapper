/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 20:31:31
 */

package open

import "medium-server-go/framework/db"

// 入驻
type Open struct {
	db.Model                      // 通用参数
	UserId         string         `json:"userId"`         // 用户 id
	ApprovalStatus ApprovalStatus `json:"approvalStatus"` // 审核状态
}

// 入驻资料
type OpenSettleIn struct {
	db.Model           // 通用参数
	OpenId      string `json:"openId"`                        // 入驻 id
	Open        Open   `json:"open" gorm:"foreignKey:OpenId"` // 入驻引用
	UserId      string `json:"userId"`                        // 用户 id
	AddressId   int    `json:"addressId"`                     // 公司地址 id
	Corporation string `json:"corporation"`                   // 公司
	Contract    string `json:"contract"`                      // 联系方式
	Name        string `json:"name"`                          // 姓名
	Address     string `json:"address"`                       // 具体地址
}

// 入驻审核
type OpenApproval struct {
	db.Model                      // 通用参数
	OpenId         string         `json:"openId"`                        // 入驻 id
	Open           Open           `json:"open" gorm:"foreignKey:OpenId"` // 入驻引用
	UserId         string         `json:"userId"`                        // 用户 id
	ApprovalStatus ApprovalStatus `json:"approvalStatus"`                // 审核状态
	Reason         string         `json:"reason"`                        // 原因
}
