/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:23:18
 */

package enum

// ApprovalStatus 审核状态
type ApprovalStatus string

const (
	ApprovalWaiting ApprovalStatus = "ApprovalWaiting" // 待审核
	ApprovalRefuse  ApprovalStatus = "ApprovalRefuse"  // 审核拒绝
	ApprovalSuccess ApprovalStatus = "ApprovalSuccess" // 审核成功
)
