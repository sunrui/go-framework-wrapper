/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 21:47:31
 */

package api_open

// 提交入驻
type postOpenReq struct {
	AddressId   int    `json:"addressId" binding:"required" validate:"len=6,numeric"` // 公司地址 id
	Corporation string `json:"corporation" binding:"required"`                        // 公司
	Contract    string `json:"contract" binding:"required"`                           // 联系方式
	Name        string `json:"name" binding:"required"`                               // 姓名
	Address     string `json:"address" binding:"required"`                            // 具体地址
}

// 提交入驻结果
type postOpenRes struct {
	OpenId string `json:"openId"` // 入驻 id
}
