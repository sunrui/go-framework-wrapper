/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/07 01:41:07
 */

package sms

import (
	"fmt"
	"math/rand"
	"medium-server-go/framework/db"
	"time"
)

// GetNowDate 获取当天日期，如 2022-01-01
func GetNowDate() string {
	now := time.Now()
	date := fmt.Sprintf("%4d-%02d-%02d", now.Year(), now.Month(), now.Day())

	return date
}

// RandomCode 创建 6 位数字
func RandomCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// CountByPhoneAndDate 获取当天验证码发送次数
func CountByPhoneAndDate(phone string, date string) int64 {
	var count int64

	query := db.Mysql.Find(&Code{}, "phone = ? AND DATE(created_at) = ?", phone, date).Count(&count)
	if query.Error != nil {
		panic(query.Error.Error())
	}

	return count
}
