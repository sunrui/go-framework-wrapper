/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 21:47:31
 */

package open

// 提交入驻
type postOpenReq struct {
	AddressId   int    `json:"addressId" validate:"required,len=6,numeric"` // 公司地址 id
	Corporation string `json:"corporation" validate:"required"`             // 公司
	Contract    string `json:"contract" validate:"required"`                // 联系方式
	Name        string `json:"name" validate:"required"`                    // 姓名
	Address     string `json:"address" validate:"required"`                 // 具体地址
}

// 提交入驻结果
type postOpenRes struct {
	OpenId string `json:"openId"` // 入驻 id
}
