/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:23:18
 */

package enum

// ApprovalStatus 审核状态
type ApprovalStatus string

const (
	Waiting ApprovalStatus = "Waiting" // 待审核
	Refuse  ApprovalStatus = "Refuse"  // 审核拒绝
	Success ApprovalStatus = "Success" // 审核成功
)
