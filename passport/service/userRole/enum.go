/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-11-19 00:36:46
 */

package userRole

// Type 类型
type Type string

const (
	AdminType     Type = "Admin"     // 管理员
	OperationType Type = "Operation" // 运营
	CockpitType   Type = "Cockpit"   // 驾驶仓
)
