/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 20:31:31
 */

package settleIn

import (
	"framework/db"
	"service/core/open"
)

// OpenSettleIn 入驻资料
type OpenSettleIn struct {
	db.Model              // 通用参数
	OpenId      string    `json:"openId"`                        // 入驻 id
	Open        open.Open `json:"open" gorm:"foreignKey:OpenId"` // 入驻引用
	UserId      string    `json:"userId"`                        // 用户 id
	AddressId   int       `json:"addressId"`                     // 公司地址 id
	Corporation string    `json:"corporation"`                   // 公司
	Contract    string    `json:"contract"`                      // 联系方式
	Name        string    `json:"name"`                          // 姓名
	Address     string    `json:"address"`                       // 具体地址
}

// 初始化
func init() {
	// 创建表入驻资料
	if err := db.Mysql.AutoMigrate(&OpenSettleIn{}); err != nil {
		panic(err.Error())
	}
}
