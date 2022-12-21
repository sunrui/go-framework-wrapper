/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-12-21 15:32:37
 */

package user

import "framework/app/mysql"

// RoleType 类型
type RoleType string

const (
	AdminType     RoleType = "ADMIN"     // 管理员
	OperationType RoleType = "OPERATION" // 运营
	CockpitType   RoleType = "COCKPIT"   // 驾驶仓
)

// UserRole 用户角色
type UserRole struct {
	mysql.Model
	UserId string `json:"userId" gorm:"type:char(12);comment:用户 id"` // 用户 id
	// Type   RoleType `json:"type" gorm:"type:varchar(32), comment:类型"`  // 类型
}
